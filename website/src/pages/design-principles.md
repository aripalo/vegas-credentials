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

