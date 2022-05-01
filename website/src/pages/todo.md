
There are already a bazillion ways to assume an IAM Role with MFA, but most existing open source tools in this scene either:
- export the temporary session credentials to environment variables (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_SESSION_TOKEN`)
- write new sections for “short-term” credentials into `~/.aws/credentials` (for example like [`aws-mfa`](https://github.com/broamski/aws-mfa) does)

The downside with those approaches is that using most of these tools (especially the ones that export environment variables) means you lose the built-in ability to automatically refresh the temporary credentials and/or the temporary credentials are “cached” in some custom location or saved into `~/.aws/credentials`.

This tool follows the concept that [you should never put temporary credentials into `~/.aws/credentials`](https://ben11kehoe.medium.com/never-put-aws-temporary-credentials-in-env-vars-or-credentials-files-theres-a-better-way-25ec45b4d73e) and also provides a mechanism to automatically refresh session credentials (if the AWS tool you use supports that).

Most AWS provided tools & SDKs already support MFA & assuming a role out of the box, but what they lack is a nice integration with Yubikey Touch, requiring you to manually type in or copy-paste the MFA TOTP token code: This utility instead integrates with [`ykman` Yubikey CLI](https://developers.yubico.com/yubikey-manager/) so just a quick touch is enough!

Also with this tool, even if you use Yubikey Touch, you still get the possibility to input MFA TOTP token code manually from an Authenticator App (for example if you don't have your Yubikey on your person).

Then there's tools such as AWS CDK that [does not support caching of assumed temporary credentials](https://github.com/aws/aws-cdk/issues/10867), requiring the user to input the MFA TOTP token code for every operation with `cdk` CLI – which makes the developer experience really cumbersome.

To recap, most existing solutions (I've seen so far) to these challenges either lack support for automatic temporary session credential refreshing, cache/write temporary session credentials to suboptimal locations and/or don't work that well with AWS tooling (i.e. requiring one to create “wrappers”):

This `vegas-credentials` is _yet another tool_, but it plugs into the standard [`credential_process`](https://docs.aws.amazon.com/sdkref/latest/guide/setting-global-credential_process.html) AWS configuration so most of AWS tooling (CLI v2, SDKs and CDK) will work out-of-the-box with it and also support automatic temporary session credential refreshing.



### Cache Mechanism

This tool caches temporary session credentials to disk. In the background it uses [`dgraph-io/badger`](https://github.com/dgraph-io/badger/) which is a fast SSD-optimized key-value store for Go and importantly supports Time-to-Live attributes for data (useful for temporary session credential expiration).

The data is stored into the key-value store with `AES-256-CTR` encryption. The encryption secret is derived from the environment (system boot time, hostname and user UID). So this is not a 100% secure setup, but it's slightly “better” compared to what [`broamski/aws-mfa`](https://github.com/broamski/aws-mfa) does or how AWS CLI caches to `~/.aws/cli/cache`: At least it provides “security by obscurity” solution against rogue scripts that might try to steal your credentials. And then again, it's only caching **_temporary_** session credentials – which you should aim to keep short-lived!

Reason why it caches temporary session credentials in the first place is to create a better user experience with AWS tools that don't support temporary session credential caching with assumed roles.

By default, the cached data is invalidated if the environment changes (system boot time, hostname and user UID), but this is okay since this tool will then query STS for new temporary session credentials and add them to cache.

Also this tool invalidates cached credentials if their expiration time is within 10 minutes and retrieves new ones. This functionality matches to [`botocore`](https://github.com/boto/botocore/blob/221ffa67a567df99ee78d7ae308c0e12d7eeeea7/botocore/credentials.py#L350-L355). You may disable this with `--disable-mandatory-refresh` CLI flag.
