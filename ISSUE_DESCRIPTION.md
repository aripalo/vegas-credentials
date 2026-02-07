# Cache decryption fails on WSL2 due to unstable boot time, causing repeated MFA prompts

## Summary

On WSL2, the credential cache becomes undecryptable after some time, causing vegas-credentials to prompt for MFA on every `credential_process` invocation instead of returning cached credentials. This is caused by `gopsutil`'s `host.BootTime()` returning unstable values on WSL2, which changes the AES encryption key used for the cache.

## Environment

- **OS:** WSL2 (`Linux 5.15.167.4-microsoft-standard-WSL2`)
- **Credential flow:** `credential_process = vegas-credentials assume --profile=<profile>` with cross-account `AssumeRole` + MFA

## Symptoms

- MFA is prompted every ~1 hour instead of being cached for the full credential duration
- Verbose output shows: `Credentials: Cache: Deserialization Error: invalid character '\u009e' looking for beginning of value`
- Credentials are successfully fetched from AWS STS with the correct duration (12h), but cannot be read back from cache on subsequent invocations

## Debugging / Root Cause Analysis

### Ruled out: AWS-side issues

1. **Role MaxSessionDuration** is correctly set to 43200 (12h):
   ```
   aws iam get-role --role-name OrganizationAccountAccessRole --profile finnkino-aider --query 'Role.MaxSessionDuration'
   43200
   ```

2. **Trust policy** has no Condition blocks restricting duration:
   ```json
   {
     "Version": "2012-10-17",
     "Statement": [{
       "Effect": "Allow",
       "Principal": { "AWS": "arn:aws:iam::<account>:root" },
       "Action": "sts:AssumeRole"
     }]
   }
   ```

3. **No restrictive SCPs** — only the default `FullAWSAccess` policy exists in the organization.

4. **Manual `aws sts assume-role --duration-seconds 43200` returns 12h credentials** — confirming AWS honors the request.

### Ruled out: vegas-credentials not passing duration

- `vegas-credentials config show-profile` correctly shows `DurationSeconds: 43200`
- The code at `internal/credentials/credentials.go:100-102` correctly passes `DurationSeconds` to the `AssumeRoleInput`
- Verbose output confirms: `Credentials: New from STS, expiration in 11 hours from now`

### Ruled out: AWS SDK for JS / CDK refresh behavior

While CDK's AWS SDK v3 is known to re-invoke `credential_process` more aggressively than the credential expiration warrants, this should be harmless — vegas-credentials should return cached credentials without prompting for MFA. The problem is that the cache read fails.

### Root cause: Unstable encryption key on WSL2

The credential cache is encrypted with AES-256-CTR. The encryption key is derived from three values (`internal/cache/encryption/encryption.go:98`):

```go
control := buildControlValue(hostname, user.Uid, bootTime)
```

The `bootTime` component is obtained via `gopsutil`'s `host.BootTime()` and formatted with **minute precision**:

```go
const bootedAtTimestampFormat string = "200601021504"
```

On WSL2, `host.BootTime()` (which reads from `/proc/stat` `btime`) returns values that drift over time. When the boot time resolves to a different minute than when the cache was written, the derived encryption key changes, and decryption of the cached credentials fails — producing the observed error:

```
Deserialization Error: invalid character '\u009e' looking for beginning of value
```

The decryption silently produces garbage bytes (AES-CTR with a wrong key doesn't error — it just produces wrong plaintext), which then fails JSON deserialization.

### Concrete evidence: `btime` drift observed on WSL2

Repeated runs of `grep btime /proc/stat` alongside `vegas-credentials assume` show the drift clearly:

| Time (EET) | `btime` | Drift from start | Cache result |
|---|---|---|---|
| 13:35:12 | 1770444586 | — | hit |
| 13:35:45 | 1770444586 | 0s | hit |
| 13:39:29 | 1770444586 | 0s | hit |
| 13:42:47 | 1770444586 | 0s | hit |
| **17:58:11** | **1770444400** | **-186s** | **MISS (deserialization error)** |
| 17:58:43 | 1770444399 | -187s | hit (new creds cached) |
| 17:58:54 | 1770444398 | -188s | hit |
| 17:58:57 | 1770444398 | -188s | hit |
| 18:00:58 | 1770444396 | -190s | hit |

Key observations:

- `btime` drifted by **186 seconds** over ~4.3 hours (~0.7 seconds/minute)
- At this drift rate, the minute-precision formatted timestamp (`200601021504`) crosses a minute boundary approximately **every 60-90 minutes**
- This perfectly matches the observed symptom of MFA prompts every ~1 hour
- After the cache miss, fresh credentials are cached with the new `btime`. Subsequent calls within the next few minutes see a `btime` that still rounds to the same minute, so the cache works — until the next minute boundary crossing

### Why this appears as "MFA every ~1 hour"

There are two caching layers at play:

1. **AWS SDK for JS (inside CDK)** — caches `credential_process` output in memory and re-invokes roughly every ~1 hour
2. **Vegas-credentials' BadgerDB cache** — persists credentials to disk, encrypted with the boot-time-derived key

When the SDK re-invokes `credential_process` after ~1 hour, the `btime` has drifted enough to cross a minute boundary, the encryption key differs, and decryption fails — triggering a new MFA prompt.

### Note on existing workaround

Windows already has a workaround for unreliable boot time — it uses a `getPreviousDay4AM()` fallback (`internal/cache/encryption/boottime/boottime.go:16-18`). WSL2 is not covered by this workaround since `runtime.GOOS` returns `"linux"` on WSL2.

## Suggested Fix

Apply the same `getPreviousDay4AM()` fallback for WSL2 that already exists for Windows. WSL2 can be detected by checking for "microsoft" or "WSL" in the kernel version string (e.g. via `/proc/version` or `uname -r`).

Alternatively, consider removing boot time from the encryption key entirely, since the cache already has TTL-based expiration via BadgerDB.
