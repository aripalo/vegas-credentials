# `aws-mfa-credential-process`
<!-- Badges -->
[![build](https://github.com/aripalo/aws-mfa-credential-process/actions/workflows/pipeline.yml/badge.svg)](https://github.com/aripalo/aws-mfa-credential-process/actions/workflows/pipeline.yml)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=aripalo_aws-mfa-credential-process&metric=sqale_rating&token=8fef16e73f5087f4331ba8b4ac6993ea8eea0e34)](https://sonarcloud.io/summary/new_code?id=aripalo_aws-mfa-credential-process)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=aripalo_aws-mfa-credential-process&metric=security_rating&token=8fef16e73f5087f4331ba8b4ac6993ea8eea0e34)](https://sonarcloud.io/summary/new_code?id=aripalo_aws-mfa-credential-process)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=aripalo_aws-mfa-credential-process&metric=vulnerabilities&token=8fef16e73f5087f4331ba8b4ac6993ea8eea0e34)](https://sonarcloud.io/summary/new_code?id=aripalo_aws-mfa-credential-process)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=aripalo_aws-mfa-credential-process&metric=coverage&token=8fef16e73f5087f4331ba8b4ac6993ea8eea0e34)](https://sonarcloud.io/summary/new_code?id=aripalo_aws-mfa-credential-process)
<!-- /Badges -->

| 🚧 🚧 🚧 <br/> **Work-in-Progress** | 🚀 ⁉️<br/> Publish Plan |
| :--------------------- | :--- |
| _**Do not use just yet (for anything real)!** Things may not work. The API and configurations may change without any prior notice at any given version. This tool is currently under development & testing. So do not use this for anything important, but feel free to test this out and give feedback!_ | After some testing (with various platforms and use case combinations), depending on the amount of bugs/issues/feedback, I'm hoping to release `v1.0.0` during November 2021 hopefully. No commitments though! |
---


<br/><br/>

[AWS `credential_process`](https://docs.aws.amazon.com/sdkref/latest/guide/setting-global-credential_process.html) utility to assume AWS IAM Roles with _[Yubikey Touch](https://www.yubico.com/products/yubikey-5-overview/) and Authenticator App_ [TOPT MFA](https://en.wikipedia.org/wiki/Time-based_One-Time_Password) to provide temporary session credentials; With encrypted caching and support for automatic credential refresh.

<br/>




![diagram](/docs/diagram.svg)

<br/>

 | [Get Started](#getting-started) | [Configuration Options](#configuration) | [Yubikey Setup](#yubikey-setup) | [Examples](#examples) | [Caveats](#caveats) | [Why yet another tool?](#why-yet-another-tool-for-this) |  [Alternatives](#alternatives) | [Cache Mechanism](#cache-mechanism) |
| :---: | :---: | :---: | :---: | :---: | :---: |  :---: | :---: |

<br/>

## Features

- **Based on AWS `credential_process`**: If you're unfamiliar with AWS `credential_process`, [this AWS re:Invent video](https://www.youtube.com/watch?v=W8IyScUGuGI&t=1260s) explains it very well

- **Supports _automatic_ temporary session credentials refreshing** for tools that understand session credential expiration

- **Supports [Role Chaining](#role-chaining)**

- **Works out-of-the-box with most AWS tools** such as AWS CDK, AWS SDKs and AWS CLI:

    Tested with AWS CDK (TypeScript), AWS CLI v2, AWS NodeJS SDK (v3), AWS Boto3 and AWS Go SDK. Should probably work with other AWS SDKs as well. Hashicorp [Terraform works as well with a minor hack](#terraform)!

- **Encrypted Caching of session credentials** to speed things up & to avoid having to input MFA token code for each operation (like in CDK)

- **Supports both Yubikey Touch or Authenticator App TOPT MFA simultaneously**: 
    
    For example you can default to using Yubikey, but if don't have the Yubikey with you all the time and also have MFA codes in an Authenticator App (such as [Authy](https://authy.com/) for example) you may just type the token code via CLI; Which ever input is given first will be used.

- **Just tap your physical key**:

    No need to manually type or copy-paste TOPT MFA token code from Yubikey Authenticator.

- **Fast & Cross-Platform**: Built with [Go](https://golang.org) and supports **macOS, Linux and Windows** operating systems with `amd64` (i.e. `x86_64`) & `arm64` (for example Apple Silicon such as M1) architectures

<br/>


## Getting Started

1. [Install `ykman`](https://developers.yubico.com/yubikey-manager/) (if you choose to use Yubikeys)

2. Install `aws-mfa-credential-process` via one of the following:

    <details><summary><strong>Homebrew</strong> (MacOS/Linux)</summary><br/>
        
    - Requires [`brew`-command](https://brew.sh/#install)
    - Install:
        ```shell
        brew tap aripalo/tap
        brew install aws-mfa-credential-process

        # Verify installation
        aws-mfa-credential-process --version
        ```
    <br/> 
    </details>

    <details><summary><strong>Scoop</strong> (Windows)</summary><br/>
    
    - Requires [`scoop`-command](https://scoop.sh#installs-in-seconds)
    - Install:
        ```shell
        scoop bucket add aripalo https://github.com/aripalo/scoops.git
        scoop install aripalo/aws-mfa-credential-process

        # Verify installation
        aws-mfa-credential-process --version
        ```
    <br/>
    </details>
    
    <details><summary><strong>NPM</strong> (MacOS/Linux/Windows)</summary><br/>

    - Requires [`node`-command](https://nodejs.org) (`v14+`)
    - Install:
        ```shell
        npm install -g aws-mfa-credential-process

        # Verify installation
        aws-mfa-credential-process --version
        ```
    <br/>
    </details>

    <details><summary><strong>Go</strong> (MacOS/Linux/Windows)</summary><br/>
    
    - Requires [`Go`-command](https://golang.org/) (`v1.17+`)
    - Perform the installation outside of any Go Module
    - Install:
        ```shell
        go install github.com/aripalo/aws-mfa-credential-process

        # Verify installation
        aws-mfa-credential-process --version
        ```
    <br/>
    </details>

    <br/>

3. Configure you source profile and its credentials, most often it's the `default` one which you configure into `~/.aws/credentials`:

    ```ini
    # ~/.aws/credentials
    [default]
    aws_access_key_id = AKIAIOSFODNN7EXAMPLE
    aws_secret_access_key = wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
    aws_mfa_device = arn:aws:iam::111111111111:mfa/MyUser
    ```

4. Configure your target profile with `credential_process` into `~/.aws/config`:

    ```ini
    # ~/.aws/config
    [profile my-profile]
    credential_process = aws-mfa-credential-process assume --profile=my-profile   
    _role_arn=arn:aws:iam::222222222222:role/MyRole # Note the underscore prefix (_)   
    _source_profile=default 
    _mfa_serial=arn:aws:iam::111111111111:mfa/MyUser
    _yubikey_serial=12345678 # if you use Yubikey, omit if you don't
    ```

5. Use any AWS tooling that support ini-based configuration with `credential_process`, like AWS CLI v2:
    ```shell
    aws sts get-caller-identity --profile my-profile
    ```

<br/>

## Configuration

There are multiple ways to configure this tool. The configuration options are evaluated in the following precedence/priority:
1. [Command-line Flag](#command-line-flags)
2. [Profile Configuration Option](#profile-configuration)
3. [Environment Variable](#environment-variables)
4. [Global Default Configuration Option](#global-defaults)


<br/>

### Command-line Flags

|       Flag        |                                                                                  Description                                                                                  |
| :---------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `--help`            | Prints help text                                                                                                                                                              |
| `--profile`         | **Required:** Which AWS Profile to use from `~/.aws/config`: Value (for example `my-profile`) must match the profile name in `ini`-section title, e.g. `[profile my-profile]` |
| `--disable-dialog`  | Disable GUI Dialog Prompt and use CLI stdin input instead                                                                                                                     |
| `--disable-mandatory-refresh` | Disable Session Credentials refreshing if expiration within 10 minutes (as defined in Botocore)                                                                                                               |
| `--hide-arns`       | Hide IAM Role & MFA Serial ARNS from output (even on verbose mode)                                                                                                            |
| `--verbose`         | Verbose output                                                                                                                                                                |
| `--debug`           | Prints out various debugging information                                                                                                                                      |
| `--no-color`        | Disable colorful fancy output                                                                                                                                                 |


<br/>

### Profile Configuration


Configuration for this tool mostly happens `~/.aws/config` ini-file. 

Important: Do not configure `role_arn`, instead provide `_role_arn`: Otherwise AWS tools would ignore the `credential_process` and assume the role directly without using this tool.

You should also prefix `source_profile` with underscore, i.e. `_source_profile`, to [avoid errors with Terraform](#terraform).


#### Standard AWS options

|       Option        |                                                                                                                                     Description                                                                                                                                      |
| :------------------ | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `credential_process`  | **Required:** To enable this tool, set the value as `aws-mfa-credential-process assume --profile <my-profile>`. Value of `my-profile` must match the profile name in `ini`-section title, e.g. `[profile my-profile]`.  |
| `source_profile`  | **Required:** Which credentials (profile) are to be used as a source for assuming the target role.                                                                                                                                                                                      |
| `mfa_serial`      | **Required:** The ARN of the Virtual (OATH TOPT) MFA device used in Multi-Factor Authentication.                                                                                                                                                                                              

You may also provide other standard [AWS options](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html#cli-configure-files-global) in `~/.aws/config`, such as `region`, `duration_seconds`, `role_session_name`, etc.

#### Custom options

|       Option        |                                                                                                                                     Description                                                                                                                                      |
| :------------------ | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `_role_arn`        | **Required:** The target IAM Role ARN to be assumed.                                                                                                                                                                                                                                         |
| `yubikey_serial`         | **Required if using Yubikey:** Yubikey Device Serial to use. You can see the serial(s) with `ykman list` command. This enforces the use of a specific Yubikey and also enables the support for using multiple Yubikeys (for different profiles)! |
| `yubikey_label`         | **Required if using any other label than the AWS MFA Device ARN as label!** Yubikey `oath` Account Label to use. You can see the available accounts with `ykman oath accounts list` command. Set the account label which you have configured your AWS TOPT MFA! |


<br/>

### Environment Variables


|   Option   |          Description          |
| :--------- | :---------------------------- |
| [`NO_COLOR`](https://no-color.org/) | Disable colorful fancy output, see also [`--no-color` CLI flag](#command-line-flags) |
| `AWS_MFA_CREDENTIAL_PROCESS_NO_COLOR` | Disable color only for this tool (not for your whole environment ) |
| `TERM=dumb` | Another way to disable colorful fancy output |

<br/>

### Global Defaults

You may define global defaults (applied to every profile) in global configuration file.

The configuration file may be written in TOML, YAML, JSON or INI with a basename of `config`, for example `config.yaml`. The configuration file is looked up from following locations in this order:
1. `$XDG_CONFIG_HOME/aws-mfa-credential-process/config.{ext}` (per [XDG-spec](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html))
2. `~/.config/aws-mfa-credential-process/config.{ext}`
3. `~/.aws-mfa-credential-process/config.{ext}`

... where `{ext}` is one of: `toml`, `yaml`, `json` or `ini`.

You may provide any [Profile Configuration option](#profile-configuration) or any behavioural (boolean) [Command-Line Flag](#command-line-flags) (i.e. all except `profile` or `help`) in this config file.

If you use a single Yubikey, then to avoid typing `yubikey_serial=12345678` into each profile in `~/.aws/config`, you should configure Yubikey Device Serial into global configuration:
```toml
# ~/.aws-mfa-credential-process/config.toml
yubikey_serial = "12345678"
```


<br/>

## Yubikey Setup

- Supports [Yubikey Touch devices](https://www.yubico.com/products/yubikey-5-overview/) with [OATH TOPT](https://en.wikipedia.org/wiki/Time-based_One-Time_Password) support (Yubikey 5 or 5C recommended).
- Requires [`ykman` CLI](https://developers.yubico.com/yubikey-manager/).
- Yubikey must be set up as [**`Virtual MFA device`** in AWS IAM](https://aws.amazon.com/blogs/security/enhance-programmatic-access-for-iam-users-using-yubikey-for-multi-factor-authentication/) - Not ~~`U2F MFA`~~!
- Think of backup strategy in case you lose your Yubikey device, you should do at least one of the following:
    - During `Virtual MFA device` setup also assign Authenticator App such as [Authy](https://authy.com/) for backup
    - If you own multiple Yubikey devices, during `Virtual MFA device` setup also configure the second Yubikey and once done, store it securely
    - Print the QR-code (and store & lock it very securely) 
    - Save the QR-code or secret key in some secure & encrypted location
- During setup, it's recommended to use `arn:aws:iam::<ACCOUNT_ID>:mfa/<IAM_USERNAME>` (i.e. MFA device ARN) as the Yubikey OATH account label:
    ```shell
    ykman oath accounts add -t arn:aws:iam::<ACCOUNT_ID>:mfa/<IAM_USERNAME>
    ```

<br/>

## Examples

### Role Chaining

This tool also supports role chaining - **given that the specific AWS tool your using supports it** - which means assuming an initial role and then using it to assume another role. An example with 3 different AWS accounts would look like:

![role-chaining](/docs/role-chaining.svg)

<br/>

Assuming correct IAM roles exists with valid permissions and trust policies:

1. Configure your long-term source credentials:
    ```ini
    # ~/.aws/credentials
    [default]
    aws_access_key_id = AKIAIOSFODNN7EXAMPLE
    aws_secret_access_key = wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
    aws_mfa_device = arn:aws:iam::111111111111:mfa/MyUser
    ```

2. Configure your profile with `credential_process`:
    ```ini
    # ~/.aws/config
    [profile my-profile]
    credential_process = aws-mfa-credential-process assume --profile=my-profile   
    _role_arn=arn:aws:iam::222222222222:role/MyRole    
    _source_profile=default 
    _mfa_serial=arn:aws:iam::111111111111:mfa/MyUser
    ```

3. Configure the final role with a `source_profile`:
    ```ini
    # ~/.aws/config
    [profile final]
    role_arn = arn:aws:iam::333333333333:role/FinalRole
    source_profile = my-profile
    ```

4. Run:
    ```shell
    aws sts get-caller-identity --profile final
    ```

<br/>

### Terraform

With Terraform AWS provider, you can not provide `source_profile` option in your profile configuration within `~/.aws/config`, if you do, Terraform will fail with an error:
```
Error: error configuring Terraform AWS Provider: Error creating AWS session: CredentialRequiresARNError: credential type source_profile requires role_arn, profile my-profile
```

To circumvent this, you may prefix this (and any other) option in `~/.aws/config` profile - that uses this credential process utility - with underscore character `_`:
```ini
# ~/.aws/config
[profile my-profile]
credential_process = aws-mfa-credential-process assume --profile=my-profile   
_role_arn=arn:aws:iam::222222222222:role/MyRole    
_source_profile=default # NOTE the underscore (_) prefix!
mfa_serial=arn:aws:iam::111111111111:mfa/MyUser
```

After that, you should be able to use Terraform as usual:
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

provider "aws" {
  profile = "my-profile"
  region  = "eu-west-1"
}
```

```sh
terraform plan
```

<br/>

## Caveats

- Does not work with [AWS SSO](https://aws.amazon.com/single-sign-on/): 

    This is by design, for AWS SSO you should use the [native SSO features](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sso.html) and fallback to using [`benkehoe/aws-sso-util`](https://github.com/benkehoe/aws-sso-util/) via `credential_process` for tooling that don't support native AWS SSO

- May work with Windows, but not tested (at least yet)

- May not output debug/info messages on some systems if `/dev/tty` not available due to [`botocore` not connecting to subprocess `stderr` on `credential_process `](https://github.com/boto/botocore/issues/1348)


<br/>


## Why yet another tool for this?

There are already a bazillion ways to assume an IAM Role with MFA, but most existing open source tools in this scene either:
- export the temporary session credentials to environment variables (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_SESSION_TOKEN`)
- write new sections for “short-term” credentials into `~/.aws/credentials` (for example like [`aws-mfa`](https://github.com/broamski/aws-mfa) does)

The downside with those approaches is that using most of these tools (especially the ones that export environment variables) means you lose the built-in ability to automatically refresh the temporary credentials and/or the temporary credentials are “cached” in some custom location or saved into `~/.aws/credentials`.

This tool follows the concept that [you should never put temporary credentials into `~/.aws/credentials`](https://ben11kehoe.medium.com/never-put-aws-temporary-credentials-in-env-vars-or-credentials-files-theres-a-better-way-25ec45b4d73e).

Most AWS provided tools & SDKs already support MFA & assuming a role out of the box, but what they lack is a nice integration with Yubikey Touch, requiring you to manually type in or copy-paste the MFA TOPT token code: This utility instead integrates with [`ykman` Yubikey CLI](https://developers.yubico.com/yubikey-manager/) so just a quick touch is enough!

Also with this tool, even if you use Yubikey Touch, you still get the possibility to input MFA TOPT token code manually from an Authenticator App (for example if you don't have your Yubikey on your person).

Then there's tools such as AWS CDK that [does not support caching of assumed temporary credentials](https://github.com/aws/aws-cdk/issues/10867), requiring the user to input the MFA TOPT token code for every operation with `cdk` CLI – which makes the developer experience really cumbersome.

To recap, most existing solutions (I've seen so far) to these challenges either lack support for automatic temporary session credential refreshing, cache/write temporary session credentials to suboptimal locations and/or don't work that well with AWS tooling (i.e. requiring one to create “wrappers”):

This `aws-mfa-credential-process` is _yet another tool_, but it plugs into the standard [`credential_process`](https://docs.aws.amazon.com/sdkref/latest/guide/setting-global-credential_process.html) AWS configuration so most of AWS tooling (CLI v2, SDKs and CDK) will work out-of-the-box with it and also support automatic temporary session credential refreshing.



<br/>

## Alternatives

- DIY shell scripts (is what people seem to often do)
- [`broamski/aws-mfa`](https://github.com/broamski/aws-mfa)
- [`99designs/aws-vault`](https://github.com/99designs/aws-vault)
- [`meeuw/aws-credential-process`](https://github.com/meeuw/aws-credential-process)
- Many others...


### Comparison

There are many great existing solutions out there that solve similar problems and I've tried to learn from them as much as I can. This tool that I've built is definitely not better than for example [`99designs/aws-vault`](https://github.com/99designs/aws-vault) in many scenarios as it has a lot more features, more contributors and been around some time. Instead the comparison below focuses on the specific use case this tool tries to solve (i.e. providing a nice UX for assuming a role with MFA using `credential_process` to support as many AWS tools as possible without having to use wrapper scripts).

|                   Feature/Info                   |    `aripalo/aws-mfa-credential-process`     |                 [`99designs/aws-vault`](https://github.com/99designs/aws-vault)                 |                  [`broamski/aws-mfa`](https://github.com/broamski/aws-mfa)                   |            [`meeuw/aws-credential-process`](https://github.com/meeuw/aws-credential-process)             |
| :----------------------------------------------- | :-----------------------------------------: | :---------------------------------------------------------------------------------------------: | :------------------------------------------------------------------------------------------: | :------------------------------------------------------------------------------------------------------: |
| Gihtub Info                                      |                   `TODO`                    | ![GitHub Repo stars](https://img.shields.io/github/stars/99designs/aws-vault?style=flat) <br/> ![GitHub last commit](https://img.shields.io/github/last-commit/99designs/aws-vault) | ![GitHub Repo stars](https://img.shields.io/github/stars/broamski/aws-mfa?style=flat)  <br/> ![GitHub last commit](https://img.shields.io/github/last-commit/broamski/aws-mfa?style=flat) | ![GitHub Repo stars](https://img.shields.io/github/stars/meeuw/aws-credential-process?style=flat) <br/> ![GitHub last commit](https://img.shields.io/github/last-commit/meeuw/aws-credential-process?style=flat) |
| `credential_process` <br/>with MFA + Assume Role |                      ✅                      |                                   ❌ [<sup>[*2]</sup>](#note2)                                   |                                 ❌ [<sup>[*4]</sup>](#note4)                                  |                                                    ✅                                                     |
| Automatic Temporary Session Credential Refresh   |                      ✅                      |                                   ❌ [<sup>[*3]</sup>](#note3)                                   |                                 ❌ [<sup>[*5]</sup>](#note5)                                  |                                                    ✅                                                     |
| Yubikey                                          |        ✅ ✅ [<sup>[*1]</sup>](#note1)        |                                   ✅ [<sup>[*1]</sup>](#note1)                                   |                                 ❌  [<sup>[*6]</sup>](#note6)                                 |                                      ✅ [<sup>[*10]</sup>](#note10)                                       |
| Cache Encryption                                 |                      ✅                      |                                                ✅                                                |                                 ❌  [<sup>[*7]</sup>](#note7)                                 |                                                    ✅                                                     |
| Cache Invalidation on config change              |                      ✅                      |                                              ✅  ?                                               |                                 ✅  [<sup>[*8]</sup>](#note8)                                 |                                                    ✅                                                     |
| Cached Performance                               | ⚡️ <br/>`<100ms`[<sup>[*11]</sup>](#note11) |                                         ⚡️ <br/>`<50ms`                                         |                              ⚡️ <br/> [<sup>[*9]</sup>](#note9)                              |                                🐢<br/>`>400ms`[<sup>[*11]</sup>](#note11)                                 |
| Comprehensively Unit Tested                      |                      ✅                      |                                                ?                                                |                                              ❌                                               |                                                    ✅                                                     |
| Installation methods                             |       `brew`, `scoop`, `npm`  , `go`        |         `brew`, `port`, `choco`, `scoop`, `pacman`, `pkg`, `zypper`, `nix-env`, `asdf`          |                                            `pip`                                             |                                              `brew`, `pip`                                               |

Please, [correct me if I'm wrong](https://github.com/aripalo/aws-mfa-credential-process/discussions) above or there's any other good alternatives!

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

## Cache Mechanism

This tool caches temporary session credentials to disk. In the background it uses [`dgraph-io/badger`](https://github.com/dgraph-io/badger/) which is a fast SSD-optimized key-value store for Go and importantly supports Time-to-Live attributes for data (useful for temporary session credential expiration).

The data is stored into the key-value store with `AES-256-CTR` encryption. The encryption secret is derived from the environment (system boot time, hostname and user UID). So this is not a 100% secure setup, but it's slightly “better” compared to what [`broamski/aws-mfa`](https://github.com/broamski/aws-mfa) does or how AWS CLI caches to `~/.aws/cli/cache`: At least it provides “security by obscurity” solution against rogue scripts that might try to steal your credentials. And then again, it's only caching **_temporary_** session credentials – which you should aim to keep short-lived!

Reason why it caches temporary session credentials in the first place is to create a better user experience with AWS tools that don't support temporary session credential caching with assumed roles.

By default, the cached data is invalidated if the environment changes (system boot time, hostname and user UID), but this is okay since this tool will then query STS for new temporary session credentials and add them to cache.

Also this tool invalidates cached credentials if their expiration time is within 10 minutes and retrieves new ones. This functionality matches to [`botocore`](https://github.com/boto/botocore/blob/221ffa67a567df99ee78d7ae308c0e12d7eeeea7/botocore/credentials.py#L350-L355). You may disable this with `--disable-mandatory-refresh` CLI flag.


<br/>

## TODO

- Error handler (At the end)
- Test manually CDK, CLI, NodeJS SDK v3, Boto3, Go ... for refresh/cache support!
- `delete-cache` command
- Add Unit tests
- Add disclaimer for orgs using this tool ("software provided as is")
- Blog post
- Contribution guidelines
- Add video that showcases the features (with CDK)
- Comparison to other solutions
- Development Docs???
- Linux & Windows support is essential
- Document auto refresh with supporting aws tools (improvement over broamski)
- Document advisory & mandatory refresh (that matches botocore)
- Document botocore retry (if less than 15*60s expiration)
- Vagrant testing / debugging for Linux & Windows ???
- Docs to https://pkg.go.dev/
- https://github.com/uber-go/guide/blob/master/style.md
- CACHE SECURITY!!
- Expiration outputs with humanize!
- See when tagged release of https://github.com/shirou/gopsutil/pull/1157 
- Github preview image (twitter etc)




<br/>

## Development

1. Fork this repository
2. Do what you need to do
3. Run all the unit tests:

    ```shell
    make test
    ```
4. Run build:

    ```shell
    make build
    ```
5. Perform manual testing
6. Push your changes
7. Create a Pull Request