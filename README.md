| **🚧 Work-in-Progress** | 🚀 ⁉️ Publish Plan |
| :--------------------- | :--- |
| Since version `v0.13.x` this tool is considered "mostly stable": Breaking changes may occur if absolutely needed but not arbitrarily. | I'm hoping to release `v1.0.0` during May 2022. No commitments though! See [`v1` Roadmap](https://github.com/aripalo/vegas-credentials/projects/1). |
---

# ![Vegas Credentials](/assets/vegas-credentials.svg "Vegas Credentials - AWS credential_process utility with optional Yubikey MFA support and smooth user experience to fetch, cache and refresh assumed temporary session credentials")

> _Much like spending a week in Las Vegas at AWS re:Invent,_ using multiple AWS tools (SDKs, CLI, CDK, Terraform, etc) via command-line to assume IAM roles in different accounts with Multi-Factor Authentication can be an exhausting experience: `vegas-credentials` aims to simplify the credential process! _And just like you shouldn't stay too long in Las Vegas at once,_ this tool only deals with temporary sesssion credentials.


Vegas Credentials is an utility with smooth user experience that plugs into AWS [`credential_process`](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html) to assume IAM Roles with [TOTP MFA](https://en.wikipedia.org/wiki/Time-based_One-Time_Password) (with optional [Yubikey Touch](https://www.yubico.com/products/yubikey-5-overview/) support) to fetch, cache and refresh assumed temporary session credentials.

<br/>

<!-- Badges -->
[![build](https://github.com/aripalo/vegas-credentials/actions/workflows/pipeline.yml/badge.svg)](https://github.com/aripalo/vegas-credentials/actions/workflows/pipeline.yml)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=aripalo_vegas-credentials&metric=coverage&token=983ccf9b47d7abae7857a352aa71fd52f953cd5c)](https://sonarcloud.io/summary/new_code?id=aripalo_vegas-credentials)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=aripalo_vegas-credentials&metric=sqale_rating&token=983ccf9b47d7abae7857a352aa71fd52f953cd5c)](https://sonarcloud.io/summary/new_code?id=aripalo_vegas-credentials)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=aripalo_vegas-credentials&metric=security_rating&token=983ccf9b47d7abae7857a352aa71fd52f953cd5c)](https://sonarcloud.io/summary/new_code?id=aripalo_vegas-credentials)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=aripalo_vegas-credentials&metric=vulnerabilities&token=983ccf9b47d7abae7857a352aa71fd52f953cd5c)](https://sonarcloud.io/summary/new_code?id=aripalo_vegas-credentials)
<!-- /Badges -->

---

<br/>

## Docs

**For guides, examples and full documentation, go to https://credentials.vegas.**



<br/>

## Install

**Via [Homebrew](https://docs.brew.sh/Installation)** on MacOS, GNU/Linux and Windows Subsystem for Linux (WSL):

```sh
brew install aripalo/tap/vegas-credentials
```

**Via [Scoop](https://scoop.sh/)** on Windows:

```sh
scoop bucket add aripalo https://github.com/aripalo/scoops.git && scoop install vegas-credentials
```


## Configure

1. Configure your source profile and its credentials, most often it's the `default` one which you configure into `~/.aws/credentials`:

    ```ini
    # ~/.aws/credentials
    [default]
    aws_access_key_id = AKIAIOSFODNN7EXAMPLE
    aws_secret_access_key = wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
    ```

2. Configure your source profile in config:

    ```ini
    # ~/.aws/config
    [default]
    mfa_serial = arn:aws:iam::111111111111:mfa/FrankSinatra
    ```

    Note: if your source profile is not `default`, remember to add `profile` as prefix (`profile foo`)

3. Configure your target profile with `credential_process` into `~/.aws/config`:

    ```ini
    # ~/.aws/config
    [profile frank@concerts]
    credential_process = vegas-credentials assume --profile=frank@concerts
    vegas_role_arn=arn:aws:iam::222222222222:role/SingerRole
    vegas_source_profile=default

    # You may also provide any other additional standard AWS configuration, such as:
    region = us-west-1
    duration_seconds = 4383
    role_session_name = SinatraAtTheSands
    external_id = 0093624694724
    ```

    Note: `role_arn` & `source_profile` must be prefixed with `vegas_` to prevent AWS tooling to ignore `credential_process` setting and to prevent Terraform failing.

4. Use any AWS tooling that support ini-based configuration with `credential_process`, like AWS CLI v2:
    ```shell
    aws sts get-caller-identity --profile frank@concerts
    ```

