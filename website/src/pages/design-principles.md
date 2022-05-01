# Design Principles

## Out of Scope Features

_By design_, some features are not implemented by `vegas-credentials` to either avoid feature bloat and also to shift responsibility to more suitable tools.

In essence, this tool **DOES NOT support**:

1. AWS SSO:

    → See [`benkehoe/aws-sso-util`](https://github.com/benkehoe/aws-sso-util) for that

2. Encrypting of master/source (long-term user) credentials in `~/.aws/credentials`

    → You may implement this quite easily [with few lines of bash & `credential_process`](https://www.youtube.com/watch?v=W8IyScUGuGI&t=1260s)

    → … or use a tool such as [`99designs/aws-vault`](https://github.com/99designs/aws-vault)

3. SAML or OpenID Connect federated authentication

4. Acting as a credential source in EC2 – i.e. `vegas-credentials` is meant for local development usage


## Your Code & Workflow MUST NOT Require Customization

Having to modify your source code or having to run custom commands before executing your AWS tools is infuriating but many existing solutions require one of the following either executing a custom script beforehand, wrapping your command somehow and/or modifying your code.

Instead `vegas-credentials` works with AWS `credential_process` which is exactly aimed for sourcing credentials via external processes _in the background_ without needing any customization in your code or workflow. Only thing that is required is a bit of configuration in your `~/.aws/config`.

If your unfamiliar with `credential_process`, [this AWS re:Invent video](https://www.youtube.com/watch?v=W8IyScUGuGI&t=1260s) explains it very well.

## Don't export anything to environment

AWS credentials are kept out of the environment due to the usage of credential sourcing in the background via `credential_process` (see above).

## Never read `~/.aws/credentials`

By design `vegas-credential` command itself never reads/parses `~/.aws/credentials` file which can contain sensitive long-term credentials: It only parses the `~/.aws/config` and then delegates the authentication process to AWS Go SDK by telling it _assume this role using this profile name_, after which the AWS Go SDK will lookup the long-term credentials for that given profile name.

## Cache Mechanism

This tool caches temporary session credentials to disk. In the background it uses [`dgraph-io/badger`](https://github.com/dgraph-io/badger/) which is a fast SSD-optimized key-value store for Go and importantly supports Time-to-Live attributes for data (useful for temporary session credential expiration).

The data is stored into the key-value store with `AES-256-CTR` encryption. The encryption secret is derived from the environment (system boot time, hostname and user UID). So this is not a 100% secure setup, but it's slightly “better” compared to what [`broamski/aws-mfa`](https://github.com/broamski/aws-mfa) does or how AWS CLI caches to `~/.aws/cli/cache`: At least it provides “security by obscurity” solution against rogue scripts that might try to steal your credentials. And then again, it's only caching **_temporary_** session credentials – which you should aim to keep short-lived!

Reason why it caches temporary session credentials in the first place is to create a better user experience with AWS tools that don't support temporary session credential caching with assumed roles.

By default, the cached data is invalidated if the environment changes (system boot time, hostname and user UID), but this is okay since this tool will then query STS for new temporary session credentials and add them to cache.

Also this tool invalidates cached credentials if their expiration time is within 10 minutes and retrieves new ones. This functionality matches to [`botocore`](https://github.com/boto/botocore/blob/221ffa67a567df99ee78d7ae308c0e12d7eeeea7/botocore/credentials.py#L350-L355).
