# Contributing

Contributions are very welcome, so thank you for considering on contributing!

- General Discussions should take place in [project's Github discussions](https://github.com/aripalo/vegas-credentials/discussions).
- Found a bug? [File a new _bug report_ issue](https://github.com/aripalo/vegas-credentials/issues/new?assignees=&labels=bug&template=bug_report.md&title=).
- Have a feature request? First condiser the [goals of this project](#goals-of-this-project) if the suggested feature applies to the scope of this tool. Then create [a new _feature request_ issue](https://github.com/aripalo/vegas-credentials/issues/new?assignees=&labels=enhancement&template=feature_request.md&title=).


## Goals of this project

The goal of this project is to provide **unified & pleasant user/developer experience for assuming AWS IAM Roles with MFA (via Yubikey or Authenticator App) using AWS `credential_process` to support as many AWS tools (such as SDKs, CLI, CDK, Terraform, etc) as reasonably possible without the user/developer having to use any wrapper scripts around AWS tools.**

In some ways, I hope this tool will become obsolete and AWS themselves would unify the role assumption, temporary session credential caching & MFA experience (with added support for Yubikeys) across AWS SDKs, CLI and CDK, but I don't think it's going to happen.

The CLI itself tries to follow guidelines for [“12 Factor CLI Apps”](https://medium.com/@jdxcode/12-factor-cli-apps-dd3c227a0e46) as much as reasonably possible.

<br/>

**By design, this tool does not support:**

- AWS SSO

    → See [`benkehoe/aws-sso-util`](https://github.com/benkehoe/aws-sso-util) for that

- Encrypting of master/source (long-term user) credentials in `~/.aws/credentials`

    → You may implement this quite easily [with few lines of bash & `credential_process`](https://www.youtube.com/watch?v=W8IyScUGuGI&t=1260s)

- SAML or OpenID Connect federated authentication

<br/>

## Development

1. You must have [Go](https://golang.org) `v1.18+` installed.

2. Fork this repository and clone it

3. Install dependencies: `go get`

4. Create a new branch, e.g. `git checkout -b feature/new-cool-thing`

5. Write code

6. Add/modify tests for the new code

7. Run tests: `go test ./..`

8. Document the new feature

9. Push your changes

10. Create a pull request
