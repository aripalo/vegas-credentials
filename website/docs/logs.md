---
sidebar_position: 6
---


# Logs


In troubleshooting situtations, it can be useful to investigate `application.log`: Running any `vegas-credentials` command logs into that (rotated) log file.

The location of the log file depends on the operation systems and user configuration. You can find the folder by running `vegas-credentials config list`:
```sh
$ vegas-credentials config list
aws config: /Users/Frank/.aws/config
ykman cli: /usr/local/bin/ykman
cache dir: /Users/Frank/Library/Caches/vegas-credentials
state dir: /Users/Frank/Library/Application Support/vegas-credentials
exec dir: /usr/local/bin
```

... from which you may lookup the `state dir`, it is usually:
- `/Users/<UserName>/Library/Application\ Support/vegas-credentials` (on MacOS)
- `C:\Users\<UserName>\AppData\Local\vegas-credentials` (on Windows)
- `$HOME/.local/state` (on GNU/Linux)

Under that directory you can find the (rotated) log file `application.log`.

An example `application.log` file looks _something_ like this:

![application-log-example](/img/application-log-example.png)

## Sensitive Data

:::info sensitive data not logged
**`vegas-credentials` will never log any sensitive information into the `application.log` file.**
:::
:::info log data not shared
**Log data is only stored in your local machine: `vegas-credentials` never sends any log (or other metrics) data anywhere.** See [Privacy Policy](/docs/privacy-policy) for more information.
:::


Example of sensitive data **NOT written to logs**:
- AWS long-term credentials (which `vegas-credentials` does not even access by itself)
- AWS short-term temporary session credentials
- Yubikey OATH application passwords

Example of data which **IS written to logs** for easier troubleshooting:
- AWS IAM Role Amazon Resource Name (ARN)
- AWS IAM MFA Virtual Device Serial ARN
- AWS Account IDs
- You local username (on the host machine)

