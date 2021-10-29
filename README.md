| 🚧 🚧 🚧  **Work-in-Progress** | 🚀 ⁉️ Publish Plan |
| :--------------------- | :--- |
| Things may break without any prior notice at any given `v0.x.x` version: So do not use this for anything critical, but feel free to test this out and give feedback! | After some testing, depending on the amount of bugs/issues/feedback, I'm hoping to release `v1.0.0` during November 2021. No commitments though! |
---

# ![Vegas Credentials](/assets/vegas-credentials.svg "Vegas Credentials - AWS credential_process utility with optional Yubikey MFA support and smooth user experience to fetch, cache and refresh assumed temporary session credentials")

> _Much like spending a week in Las Vegas at AWS re:Invent,_ using multiple AWS tools (SDKs, CLI, CDK, Terraform, etc) via command-line to assume IAM roles in different accounts with Multi-Factor Authentication can be an exhausting experience: `vegas-credentials` aims to simplify the credential process! _And just like you shouldn't stay too long in Las Vegas at once,_ this tool only deals with temporary sesssion credentials.


Vegas Credentials is an utility with smooth user experience that plugs into AWS [`credential_process`](https://docs.aws.amazon.com/sdkref/latest/guide/setting-global-credential_process.html) to assume IAM Roles with [TOPT MFA](https://en.wikipedia.org/wiki/Time-based_One-Time_Password) (with optional [Yubikey Touch](https://www.yubico.com/products/yubikey-5-overview/) support) to fetch, cache and refresh assumed temporary session credentials.

<br/>

<!-- Badges -->
[![build](https://github.com/aripalo/vegas-credentials/actions/workflows/pipeline.yml/badge.svg)](https://github.com/aripalo/vegas-credentials/actions/workflows/pipeline.yml)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=aripalo_vegas-credentials&metric=coverage&token=983ccf9b47d7abae7857a352aa71fd52f953cd5c)](https://sonarcloud.io/summary/new_code?id=aripalo_vegas-credentials)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=aripalo_vegas-credentials&metric=sqale_rating&token=983ccf9b47d7abae7857a352aa71fd52f953cd5c)](https://sonarcloud.io/summary/new_code?id=aripalo_vegas-credentials)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=aripalo_vegas-credentials&metric=security_rating&token=983ccf9b47d7abae7857a352aa71fd52f953cd5c)](https://sonarcloud.io/summary/new_code?id=aripalo_vegas-credentials)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=aripalo_vegas-credentials&metric=vulnerabilities&token=983ccf9b47d7abae7857a352aa71fd52f953cd5c)](https://sonarcloud.io/summary/new_code?id=aripalo_vegas-credentials)
<!-- /Badges -->

---

| [Features](#features) | [Overview](#overview) | [Get&nbsp;Started](#get-started) | [Yubikeys](#yubikeys) | [Configuration](#configuration) | [Examples](#examples) | [Why?](#why-yet-another-tool) | [Design](#design) |
| :-------------------: | :-------------------: | :------------------------------: | :-------------------: | :-----------------------------: | :-------------------: | :----------: | :---------------: |

---

## Features

- **Plugs into AWS `credential_process`**: If you're unfamiliar with AWS `credential_process`, [this AWS re:Invent video](https://www.youtube.com/watch?v=W8IyScUGuGI&t=1260s) explains it very well

- **Supports _automatic_ temporary session credential refreshing** for tools that understand session credential expiration

- **Supports [Role Chaining](#role-chaining)**

- **Works out-of-the-box with most tools** such as AWS CLI, [most](https://docs.aws.amazon.com/sdkref/latest/guide/setting-global-credential_process.html#setting-credential_process-sdk-compat) AWS SDKs, AWS CDK, Terraform...

- **Encrypted Caching of session credentials** to speed things up & to avoid having to input MFA token code for each operation

- **Supports both Yubikey Touch or Authenticator App TOPT MFA _simultaneously_**: 
    
    - For example you can default to using Yubikey, but if don't have the Yubikey with you all the time and also have MFA codes in an Authenticator App (such as [Authy](https://authy.com/) for example) 
    - You may just touch the Yubikey or manually type the token code (via GUI Prompt Dialog or CLI `stdin`) – which ever input is given first will be used

- **Smooth Yubikey integration**:

    - Just tap your physical key: No need to manually type or copy-paste MFA token code
    - Supports multiple Yubikey devices

- **Fast & Cross-Platform**: Built with [Go](https://golang.org) and supports **macOS, Linux and Windows** operating systems with `x86_64` & `arm64` (e.g. Apple M1) architectures

<br/>

### By design, this tool _does not_ support:

- AWS SSO 

    → See [`benkehoe/aws-sso-util`](https://github.com/benkehoe/aws-sso-util) for that

- Encrypting of master/source (long-term user) credentials in `~/.aws/credentials`

    → You may implement this quite easily [with few lines of bash & `credential_process`](https://www.youtube.com/watch?v=W8IyScUGuGI&t=1260s)
    
    → … or use a tool such as [`99designs/aws-vault`](https://github.com/99designs/aws-vault)

- SAML or OpenID Connect federated authentication

- Acting as a credential source in EC2 – i.e. `vegas-credentials` is meant for local development usage


<br/>

<br/>

## Overview

![diagram](/docs/diagram.svg)

<br/>

## Get Started

1. Install `vegas-credentials` via one of the following:

    <details><summary><strong>Homebrew</strong> (MacOS/Linux)</summary><br/>
        
    - Requires [`brew`-command](https://brew.sh/#install)
    - Install:
        ```shell
        brew tap aripalo/tap
        brew install vegas-credentials

        # Verify installation
        vegas-credentials --version
        ```
    <br/> 
    </details>

    <details><summary><strong>Scoop</strong> (Windows)</summary><br/>
    
    - Requires [`scoop`-command](https://scoop.sh#installs-in-seconds)
    - Install:
        ```shell
        scoop bucket add aripalo https://github.com/aripalo/scoops.git
        scoop install aripalo/vegas-credentials

        # Verify installation
        vegas-credentials --version
        ```
    <br/>
    </details>
    
    <details><summary><strong>NPM</strong> (MacOS/Linux/Windows)</summary><br/>

    - Requires [`node`-command](https://nodejs.org) (`v14+`)
    - Install:
        ```shell
        npm install -g vegas-credentials

        # Verify installation
        vegas-credentials --version
        ```
    <br/>
    </details>

    <br/>

2. Configure you source profile and its credentials, most often it's the `default` one which you configure into `~/.aws/credentials`:

    ```ini
    # ~/.aws/credentials
    [default]
    aws_access_key_id = AKIAIOSFODNN7EXAMPLE
    aws_secret_access_key = wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
    ```

3. TODO

    ```ini
    # ~/.aws/config
    [default]
    mfa_serial = arn:aws:iam::111111111111:mfa/FrankSinatra
    ```

4. Configure your target profile with `credential_process` into `~/.aws/config`:

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

5. Use any AWS tooling that support ini-based configuration with `credential_process`, like AWS CLI v2:
    ```shell
    aws sts get-caller-identity --profile frank@concerts
    ```


<br/>

## Yubikeys

To use Yubikeys:

1. You must have at least one [Yubikey Touch device](https://www.yubico.com/products/yubikey-5-overview/) with [OATH TOPT](https://en.wikipedia.org/wiki/Time-based_One-Time_Password) support (Yubikey 5 or 5C recommended).
2. Install [`ykman` CLI](https://developers.yubico.com/yubikey-manager/)
3. Set up Yubikey as [**`Virtual MFA device`** in AWS IAM](https://aws.amazon.com/blogs/security/enhance-programmatic-access-for-iam-users-using-yubikey-for-multi-factor-authentication/) - Not ~~`U2F MFA`~~!
4. Think of backup strategy in case you lose your Yubikey device, you should do at least one of the following:
    - During `Virtual MFA device` setup also assign Authenticator App such as [Authy](https://authy.com/) for backup
    - If you own multiple Yubikey devices, during `Virtual MFA device` setup also configure the second Yubikey and once done, store it securely
    - Print the QR-code (and store & lock it very securely) 
    - Save the QR-code or secret key in some secure & encrypted location
5. When configuring your Yubikey, use `arn:aws:iam::<ACCOUNT_ID>:mfa/<IAM_USERNAME>` (i.e. MFA device ARN) as the Yubikey OATH account label:
    ```shell
    ykman oath accounts add -t arn:aws:iam::<ACCOUNT_ID>:mfa/<IAM_USERNAME>
    ```
6. Configure `vegas_yubikey_serial` into `~/.aws/config`:
    ```ini
    # ~/.aws/config
    [default]
    mfa_serial = arn:aws:iam::111111111111:mfa/FrankSinatra
    vegas_yubikey_serial = 12345678
    ```

    Optionally you may configure `vegas_yubikey_label` if you used something else than `arn:aws:iam::<ACCOUNT_ID>:mfa/<IAM_USERNAME>` as the OATH account label (though not recommended).


<br/>

## Configuration

Configuration for the most part happens in `~/.aws/config` ini-file, but there are some command-line flags and environment variables that you may want to use sometimes.

### Source Profile Configuration

|         Option         |                                                                                                                           Description                                                                                                                           |
| :--------------------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `mfa_serial`           | **Required:** The ARN of the Virtual (OATH TOPT) MFA device used in Multi-Factor Authentication.                                                                                                                                                                |
| `vegas_yubikey_serial` | **Required if using Yubikey:** Yubikey Device Serial to use. You can see the serial(s) with `ykman list` command. This enforces the use of a specific Yubikey and also enables the support for using multiple Yubikeys (for different profiles)!                |
| `vegas_yubikey_label`  | Use only if you have any other value than the AWS MFA Device ARN as `oath` account label! Yubikey `oath` Account Label to use. You can see the available accounts with `ykman oath accounts list` command. Set the account label which you have configured your AWS TOPT MFA! |

Example:
```ini
# ~/.aws/config
[default]
mfa_serial = arn:aws:iam::111111111111:mfa/FrankSinatra
vegas_yubikey_serial = 12345678
```


### Target Profile Configuration

|         Option         |                                                                                                      Description                                                                                                       |
| :--------------------- | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `credential_process`   | **Required:** To enable this tool, set the value as `vegas-credentials assume --profile <my-profile>`. Value of `my-profile` must match the profile name in `ini`-section title, e.g. `[profile my-profile]`. |
| `vegas_role_arn`       | **Required:** The target IAM Role ARN to be assumed.                                                                                                                                                                   |
| `vegas_source_profile` | **Required:** Which credentials (profile) are to be used as a source for assuming the target role.                                                                                                                     |

Example:
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


### Command-line Flags

|       Flag        |                                                                                  Description                                                                                  |
| :---------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `--help`            | Prints help text                                                                                                                                                              |
| `--profile`         | **Required:** Which AWS Profile to use from `~/.aws/config`: Value (for example `my-profile`) must match the profile name in `ini`-section title, e.g. `[profile my-profile]` |
| `--disable-dialog`  | Disable GUI Dialog Prompt                                                                                                                   |
| `--disable-mandatory-refresh` | Disable Session Credentials refreshing if expiration within 10 minutes (as defined in Botocore)                                                                                                               |
| `--hide-arns`       | Hide IAM Role & MFA Serial ARNS from output (even on verbose mode)                                                                                                            |
| `--verbose`         | Verbose output                                                                                                                                                                |
| `--debug`           | Prints out various debugging information                                                                                                                                      |
| `--no-color`        | Disable colorful fancy output                                                                                                                                                 |

Example:
```sh
vegas-credentials assume --profile=frank@concerts --verbose --no-color
```
... though you shouldn't really call this tool directly yourself, but instead configure it as `credential_process` in `~/.aws/config`.

### Environment Variables


|   Option   |          Description          |
| :--------- | :---------------------------- |
| [`NO_COLOR`](https://no-color.org/) | Disable colorful fancy output, see also [`--no-color` CLI flag](#command-line-flags) |
| `VEGAS_CREDENTIALS_NO_COLOR` | Disable color only for this tool (not for your whole environment ) |
| `TERM=dumb` | Another way to disable colorful fancy output |



<br/>

## Examples

### Using with AWS SDKs

Often times you may not want to define the `profile` within the application code, since the application code most often will be ran without profile in cloud. You may circumvent this by setting the profile via environment variable:
```shell
AWS_PROFILE=frank@concerts ts-node src/index.ts
```

> _By default, the SDK checks the `AWS_PROFILE` environment variable to determine which profile to use. If the `AWS_PROFILE` variable is not set in your environment, the SDK uses the credentials for the `[default]` profile. To use one of the alternate profiles, set or change the value of the `AWS_PROFILE` environment variable. For example, given the configuration file shown, to use the credentials from the work account, set the `AWS_PROFILE` environment variable to work-account (as appropriate for your operating system)._

### Role Chaining

This tool also supports role chaining - **given that the specific AWS tool your using supports it** - which means assuming an initial role and then using it to assume another role. An example with 3 different AWS accounts would look like:

![role-chaining](/docs/role-chaining.svg)

<br/>

Assuming correct IAM roles exists with valid permissions and trust policies:

1. Assuming you have the following configuration already:
    ```ini
    # ~/.aws/credentials
    [default]
    aws_access_key_id = AKIAIOSFODNN7EXAMPLE
    aws_secret_access_key = wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
    ```

    ```ini
    # ~/.aws/config
    [default]
    aws_mfa_device = arn:aws:iam::111111111111:mfa/FrankSinatra
    
    [profile frank@concerts]
    credential_process = vegas-credentials assume --profile=frank@concerts
    vegas_role_arn=arn:aws:iam::222222222222:role/SingerRole    
    vegas_source_profile=default 
    ```

2. Configure another role with standard `role_arn` and `source_profile`:
    ```ini
    # ~/.aws/config
    [profile frank@movies]
    role_arn=arn:aws:iam::333333333333:role/ActorRole # Important: NO prefix here!   
    source_profile=frank@concerts # Important: NO prefix here!
    ```

3. Do some chaining:
    ```shell
    aws sts get-caller-identity --profile frank@movies
    ```

<br/>

### Terraform

```tf
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.27"
    }
  }

  required_version = ">= 0.14.9"
}

# Again, nothing special here, just normal profile configuration…
provider "aws" {
  profile = "frank@concerts"
  region  = "eu-west-3"
}

data "aws_region" "current" {}

output "album_name" {
  value = "Sinatra & Sextet: Live in ${
    replace(
      regex("[a-zA-Z]+\\)$", data.aws_region.current.description),
      ")",
      ""
    )
  }"
}
```

<br/>

## Why yet another tool?

### Reasons



There are already a bazillion ways to assume an IAM Role with MFA, but most existing open source tools in this scene either:
- export the temporary session credentials to environment variables (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_SESSION_TOKEN`)
- write new sections for “short-term” credentials into `~/.aws/credentials` (for example like [`aws-mfa`](https://github.com/broamski/aws-mfa) does)

The downside with those approaches is that using most of these tools (especially the ones that export environment variables) means you lose the built-in ability to automatically refresh the temporary credentials and/or the temporary credentials are “cached” in some custom location or saved into `~/.aws/credentials`.

This tool follows the concept that [you should never put temporary credentials into `~/.aws/credentials`](https://ben11kehoe.medium.com/never-put-aws-temporary-credentials-in-env-vars-or-credentials-files-theres-a-better-way-25ec45b4d73e) and also provides a mechanism to automatically refresh session credentials (if the AWS tool you use supports that).

Most AWS provided tools & SDKs already support MFA & assuming a role out of the box, but what they lack is a nice integration with Yubikey Touch, requiring you to manually type in or copy-paste the MFA TOPT token code: This utility instead integrates with [`ykman` Yubikey CLI](https://developers.yubico.com/yubikey-manager/) so just a quick touch is enough!

Also with this tool, even if you use Yubikey Touch, you still get the possibility to input MFA TOPT token code manually from an Authenticator App (for example if you don't have your Yubikey on your person).

Then there's tools such as AWS CDK that [does not support caching of assumed temporary credentials](https://github.com/aws/aws-cdk/issues/10867), requiring the user to input the MFA TOPT token code for every operation with `cdk` CLI – which makes the developer experience really cumbersome.

To recap, most existing solutions (I've seen so far) to these challenges either lack support for automatic temporary session credential refreshing, cache/write temporary session credentials to suboptimal locations and/or don't work that well with AWS tooling (i.e. requiring one to create “wrappers”):

This `vegas-credentials` is _yet another tool_, but it plugs into the standard [`credential_process`](https://docs.aws.amazon.com/sdkref/latest/guide/setting-global-credential_process.html) AWS configuration so most of AWS tooling (CLI v2, SDKs and CDK) will work out-of-the-box with it and also support automatic temporary session credential refreshing.


### Alternatives

There are many great existing solutions out there that solve similar problems and I've tried to learn from them as much as I can. This tool that I've built is definitely not better than for example [`99designs/aws-vault`](https://github.com/99designs/aws-vault) in many scenarios as it has a lot more features, more contributors and been around some time. Instead the comparison below focuses on the specific use case this tool tries to solve (i.e. providing a nice UX for assuming a role with MFA using `credential_process` to support as many AWS tools as possible without having to use wrapper scripts).

|                   Feature/Info                   |    `aripalo/vegas-credentials`     |                 [`99designs/aws-vault`](https://github.com/99designs/aws-vault)                 |                  [`broamski/aws-mfa`](https://github.com/broamski/aws-mfa)                   |            [`meeuw/aws-credential-process`](https://github.com/meeuw/aws-credential-process)             |
| :----------------------------------------------- | :-----------------------------------------: | :---------------------------------------------------------------------------------------------: | :------------------------------------------------------------------------------------------: | :------------------------------------------------------------------------------------------------------: |
| Github Info                                      |                   `TODO`                    | ![GitHub Repo stars](https://img.shields.io/github/stars/99designs/aws-vault?style=flat) <br/> ![GitHub last commit](https://img.shields.io/github/last-commit/99designs/aws-vault) | ![GitHub Repo stars](https://img.shields.io/github/stars/broamski/aws-mfa?style=flat)  <br/> ![GitHub last commit](https://img.shields.io/github/last-commit/broamski/aws-mfa?style=flat) | ![GitHub Repo stars](https://img.shields.io/github/stars/meeuw/aws-credential-process?style=flat) <br/> ![GitHub last commit](https://img.shields.io/github/last-commit/meeuw/aws-credential-process?style=flat) |
| `credential_process` <br/>with MFA + Assume Role |                      ✅                      |                                   ❌ [<sup>[*2]</sup>](#note2)                                   |                                 ❌ [<sup>[*4]</sup>](#note4)                                  |                                                    ✅                                                     |
| Automatic Temporary Session Credential Refresh   |                      ✅                      |                                   ❌ [<sup>[*3]</sup>](#note3)                                   |                                 ❌ [<sup>[*5]</sup>](#note5)                                  |                                                    ✅                                                     |
| Yubikey                                          |        ✅ ✅ [<sup>[*1]</sup>](#note1)        |                                   ✅ [<sup>[*1]</sup>](#note1)                                   |                                 ❌  [<sup>[*6]</sup>](#note6)                                 |                                      ✅ [<sup>[*10]</sup>](#note10)                                       |
| Cache Encryption                                 |                      ✅                      |                                                ✅                                                |                                 ❌  [<sup>[*7]</sup>](#note7)                                 |                                                    ✅                                                     |
| Cache Invalidation on config change              |                      ✅                      |                                              ✅  ?                                               |                                 ✅  [<sup>[*8]</sup>](#note8)                                 |                                                    ✅                                                     |
| Cached Performance                               | ⚡️ <br/>`<100ms`[<sup>[*11]</sup>](#note11) |                                         ⚡️ <br/>`<50ms`                                         |                              ⚡️ <br/> [<sup>[*9]</sup>](#note9)                              |                                🐢<br/>`>400ms`[<sup>[*11]</sup>](#note11)                                 |
| Comprehensively Unit Tested                      |                      ✅                      |                                                ?                                                |                                              ❌                                               |                                                    ✅                                                     |
| Installation methods                             |       `brew`, `scoop`, `npm`         |         `brew`, `port`, `choco`, `scoop`, `pacman`, `pkg`, `zypper`, `nix-env`, `asdf`          |                                            `pip`                                             |                                              `brew`, `pip`                                               |

Please, [correct me if I'm wrong](https://github.com/aripalo/vegas-credentials/discussions) above or there's any other good alternatives!

#### Notes

#### `99designs/aws-vault`

1. <a id="note1"></a>Yubikey support in `99designs/aws-vault` is not perfect:
    - Using multiple Yubikeys is cumbersome due to having to pass in Yubikey [device serial as environment variable for each command](https://github.com/99designs/aws-vault/pull/748) – vs. this tool allows setting device serial via configuration per profile (no need to remember the serial for each Yubikey).
    - Uses deprecated `ykman` commands.
    - See also [point 2](#note2) about `credential_process`, assumed roles and Yubikeys.

2. <a id="note2"></a>Does not seem to play well with `credential_process`:
    - **At least I haven't figured out how to succesfully configure it to use `credential_process`, assume a role, use Yubikey for MFA and to provide temporary session credentials.** 
    - They themselves [claim that _“`credential_process` is designed for retrieving master credentials”_](https://github.com/99designs/aws-vault/blob/master/USAGE.md#using-credential_proce) - which is NOT true since this tool does work with temporary credentials via `credential_process` just fine and even the [AWS docs on `credential_process` show `SessionToken` and `Expiration` on the expected output from the credentials program](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html).
    - There's further indication that [`99designs/aws-vault` is not designed for `credential_process`](https://github.com/99designs/aws-vault/issues/641#issuecomment-681346113):
       
        > _**Using credentials_process isn't the way I use aws-vault**, it was a contributed addition, so feels like we should emphasise this is not the recommended path._
        >
        > – Michael Tibben, VP Technology, 99designs

3. <a id="note3"></a>This pretty much relates to [point 1](#note1): For AWS tools to automatically request refreshed credentials, the credentials need to be provided via either the multiple standard methods or via `credential_process`.


#### `broamski/aws-mfa`

4. <a id="note4"></a>Works differently by writing temporary session credentials into `~/.aws/credentials`, so therefore no `credential_process` support at all.

5. <a id="note5"></a>If temporary session credentials written into `~/.aws/credentials` by `broamski/aws-mfa` are expired, AWS tools will fail and you must invoke `aws-mfa` command manually to fetch new session credentials. There is no (automatic) way for AWS tools to trigger `aws-mfa` command.

6. <a id="note6"></a>You may use Yubikey, but it requires you to manually copy-paste the value from `ykman` or Yubikey Manager GUI. No "touch integration".

7. <a id="note7"></a>Temporary session credentials are written in plaintext into `~/aws/credentials`. Besides being available as plaintext, it pollutes the credentials file.

8. <a id="note8"></a>Configuration is only provided via flags to `aws-mfa` CLI command, so each time you execute `aws-mfa` it will use the flags provided. But, the gotcha is that again you need to execute `aws-mfa` manually always.

9. <a id="note9"></a>As temporary session credentials (or "short-term" as `aws-mfa` calls them) are stored as plaintext into `~/aws/credentials`, there is no delay since AWS tools can directly read them from that file.

#### `meeuw/aws-credential-process`

10. <a id="note10"></a>Does not support multiple Yubikey devices.

11. <a id="note11"></a>Performance

    [Hyperfine](https://github.com/sharkdp/hyperfine) benchmark for retrieving cached temporary session credentials:

    ![perf](/docs/perf-comparison.png)


<br/>

## Design


### Cache Mechanism

This tool caches temporary session credentials to disk. In the background it uses [`dgraph-io/badger`](https://github.com/dgraph-io/badger/) which is a fast SSD-optimized key-value store for Go and importantly supports Time-to-Live attributes for data (useful for temporary session credential expiration).

The data is stored into the key-value store with `AES-256-CTR` encryption. The encryption secret is derived from the environment (system boot time, hostname and user UID). So this is not a 100% secure setup, but it's slightly “better” compared to what [`broamski/aws-mfa`](https://github.com/broamski/aws-mfa) does or how AWS CLI caches to `~/.aws/cli/cache`: At least it provides “security by obscurity” solution against rogue scripts that might try to steal your credentials. And then again, it's only caching **_temporary_** session credentials – which you should aim to keep short-lived!

Reason why it caches temporary session credentials in the first place is to create a better user experience with AWS tools that don't support temporary session credential caching with assumed roles.

By default, the cached data is invalidated if the environment changes (system boot time, hostname and user UID), but this is okay since this tool will then query STS for new temporary session credentials and add them to cache.

Also this tool invalidates cached credentials if their expiration time is within 10 minutes and retrieves new ones. This functionality matches to [`botocore`](https://github.com/boto/botocore/blob/221ffa67a567df99ee78d7ae308c0e12d7eeeea7/botocore/credentials.py#L350-L355). You may disable this with `--disable-mandatory-refresh` CLI flag.

### Never touch `~/.aws/credentials`

TODO

### Never export credentials to environment

TODO