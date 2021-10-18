# `aws-mfa-credential-process`

| 🚧 🚧 🚧 <br/> **Work-in-Progress** |
| :---------------------: |
| _**Do not use just yet!** Things may not work. The API and configurations may change without any prior notice at any given version. This tool is currently under development & testing. So do not use this for anything important, but feel free to test this out and give feedback!_ |



<br/><br/>

[AWS `credential_process`](https://docs.aws.amazon.com/sdkref/latest/guide/setting-global-credential_process.html) utility to assume AWS IAM Roles with _[Yubikey Touch](https://www.yubico.com/products/yubikey-5-overview/) and Authenticator App_ [TOPT MFA](https://en.wikipedia.org/wiki/Time-based_One-Time_Password) to provide temporary session credentials; With caching to local [keyring](#keyring) and support for automatic credential refresh.

If you're unfamiliar with AWS `credential_process`, [this AWS re:Invent video](https://www.youtube.com/watch?v=W8IyScUGuGI&t=1260s) explains it very well.

<br/>




![diagram](/docs/diagram.svg)

<br/>

| [Features](#features) | [Get Started](#getting-started) | [Configuration](#configuration) | [Keyring](#keyring) | [Yubikey](#yubikey-setup) |  [Why yet another tool?](#why-yet-another-tool-for-this) | [Caveats](#caveats) |[TODO](#todo) |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |

<br/>

## Features

- **Supports _automatic_ temporary session credentials refreshing** for tools that understand session credential expiration

- **Works out-of-the-box with most AWS tools** such as AWS CDK, AWS SDKs and AWS CLI:

    Tested with AWS CDK (TypeScript), AWS CLI v2, AWS NodeJS SDK (v3), AWS Boto3 and AWS Go SDK. Should probably work with other AWS SDKs as well.

- **Caching of session credentials into [Keyring](#keyring)** to speed things up & to avoid having to input MFA token code for each operation

- **Supports both Yubikey Touch or Authenticator App TOPT MFA simultaneously**: 
    
    For example you can default to using Yubikey, but if don't have the Yubikey with you all the time and also have MFA codes in an Authenticator App (such as [Authy](https://authy.com/) for example) you may just type the token code via CLI; Which ever input is given first will be used.

- **Just tap your physical key**:

    No need to manually type or copy-paste TOPT MFA token code from Yubikey Authenticator.

- **Fast & Cross-Platform**: Built with [Go](https://golang.org) and supports **macOS, Linux and Windows** operating systems with `amd64` (i.e. `x86_64`) & `arm64` (for example Apple Silicon such as M1) architectures

<br/>


## Getting Started

1. [Install `ykman`](https://developers.yubico.com/yubikey-manager/) (if you choose to use Yubikeys)

2. **TODO** Install:

    1. via [Homebrew](https://brew.sh/):
        ```shell
        brew tap aripalo/tap
        brew install aws-mfa-credential-process
        ```

    2. or manual download via Github releases:
        1. Go to https://github.com/aripalo/aws-mfa-credential-process/releases
        2. Download the latest binary for your platform & operating system
        3. Save the binary to your preferred location
        4. In some directory that is in your `$PATH`, create a symlink of `aws-mfa-credential-process` pointing into the downloaded binary

    3. more installation methods **maybe implemented later**:
        - `go install github.com/aripalo/aws-mfa-credential-process` ?
        - `pipx install aws-mfa-credential-process` ? Convenient as there's lot of AWS CLI, Boto3 and Botocore users...
        - `npm install -g @aripalo/aws-mfa-credential-process` ? Convenient as there's a lot of AWS CDK & AWS NodeJS SDK users...

3. Configure you source profile and its credentials, most often it's the `default` one which you configure into `~/.aws/credentials`:

    ```ini
    [default]
    aws_access_key_id = AKIAIOSFODNN7EXAMPLE
    aws_secret_access_key = wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
    aws_mfa_device = arn:aws:iam::123456789012:mfa/example
    ```

4. Configure your target profile with `credential_process` into `~/.aws/config`:

    ```ini
    [profile my-profile]
    credential_process = aws-mfa-credential-process assume --profile=my-profile   # in 99% cases the profile should match what's in the ini-section title
    yubikey_serial=<yubikey-serial>        # which Yubikey Device you're using
    assume_role_arn=<target-role-arn>      # IMPORTANT: Note "assume_" prefix (vs. the default "role_arn")
    source_profile=<source-profile-name>   # often "default"
    mfa_serial=<mfa-device-arn>            # your AWS MFA Device ARN serial
    # ...any other AWS standard options
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
| `--disable-refresh` | Disable Session Credentials refreshing (as defined in Botocore)                                                                                                               |
| `--hide-arns`       | Hide IAM Role & MFA Serial ARNS from output (even on verbose mode)                                                                                                            |
| `--verbose`         | Verbose output                                                                                                                                                                |
| `--debug`           | Prints out various debugging information                                                                                                                                      |
| `--no-color`        | Disable colorful fancy output                                                                                                                                                 |


<br/>

### Profile Configuration


Configuration for this tool mostly happens `~/.aws/config` ini-file. 

Important: Do not configure `role_arn`, instead provide `assume_role_arn`: Otherwise AWS tools would ignore the `credential_process` and assume the role directly without using this tool.


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
| `assume_role_arn`        | **Required:** The target IAM Role ARN to be assumed.                                                                                                                                                                                                                                         |
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





<br/>

## Keyring

In the background this tool uses [`99designs/keyring`](https://github.com/99designs/keyring) to integrate with:
- macOS/OSX Keychain
- Windows credential store
- Pass
- Secret Service
- KDE Wallet
- Encrypted File

🚧 **TODO**: Not all keyring backends are tested/verified working!

### Keyring specific notes

<details><summary>macOS Keychain</summary><br>


This tool will create a new keychain with a name `aws-mfa-credential-process`.

During first use of the tool, you must create a password for the keychain and later you must provide the password to access items in it. For convenience you should choose `Always Allow`, but see [security related notes](#security-related-notes).

| 1. Initial setup | 2. Usage |
|:---: | :---: |
| ![new-keyring](/docs/keyring/1-new-keyring.png) | ![use-keyring-info](/docs/keyring/3-use-keyring-info.png) | 

</details><br>

### Security related notes

Keyring (such as macOS Keychain) by itself probably is one of the most secure places to cache temporary session credentials, but beware: If you allow `aws-mfa-credential-process` to access the keyring, it means that any (potentially hostile) process on your computer could invoke `aws-mfa-credential-process` and retrieve the cached credentials! So it's not 100% secure. Then again, we're storing only _**temporary** session credentials_ there with an expiration and you really should use temporary credentials with short expiration (`3600s` i.e. 1 hour is a good default for expiration).


<br/>

## Yubikey Setup

- Supports [Yubikey Touch devices](https://www.yubico.com/products/yubikey-5-overview/) with [OATH TOPT](https://en.wikipedia.org/wiki/Time-based_One-Time_Password) support (Yubikey 5 or 5C recommended).
- Requires [`ykman` CLI](https://developers.yubico.com/yubikey-manager/).
- Yubikey must be set up as [**`Virtual MFA device`** in AWS IAM](https://aws.amazon.com/blogs/security/enhance-programmatic-access-for-iam-users-using-yubikey-for-multi-factor-authentication/) - Not ~~`U2F MFA`~~!
- Think of backup strategy in case you lose your Yubikey device, you should do at least one of the following:
    - Print the QR-code (and store it securely) 
    - During `Virtual MFA device` setup also assign Authenticator App (such as [Authy](https://authy.com/)) for backup in case you lose your Yubikey
    - If you own multiple Yubikey devices, during `Virtual MFA device` setup also configure the second Yubikey and once done, store it securely
- During setup, it's recommended to use `arn:aws:iam::<ACCOUNT_ID>:mfa/<IAM_USERNAME>` (i.e. MFA device ARN) as the Yubikey OATH account label.

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

## Caveats

- Does not work with [AWS SSO](https://aws.amazon.com/single-sign-on/): 

    This is by design, for AWS SSO you should use the [native SSO features](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sso.html) and fallback to using [`benkehoe/aws-sso-util`](https://github.com/benkehoe/aws-sso-util/) via `credential_process` for tooling that don't support native AWS SSO

- May work with Windows, but not tested (at least yet)

- May not output debug/info messages on some systems if `/dev/tty` not available due to [`botocore` not connecting to subprocess `stderr` on `credential_process `](https://github.com/boto/botocore/issues/1348)

<br/>

## TODO


- Ensure Role Chaining Works!
- Test manually CDK, CLI, NodeJS SDK v3, Boto3, Go ... for refresh/cache support!
- Add Unit tests
- Add disclaimer for orgs using this tool ("software provided as is")
- Add PR templates (bug, feature request...)
- Blog post
- Development docs (maybe separate site?)
- Contribution guidelines
- Add video that showcases the features (with CDK)
- Documentation pages (docusaurus to Github Pages)
    - Custom domain?
    - Most of the stuff from README
    - Examples
    - Video
    - Logging/Debugging
    - Session Credential Refresh
    - Role Chaining
    - Getting Started
    - Configuration
    - Keyring
    - Recommended Yubikeys
    - How to setup Yubikey for TOPT MFA (and additionally add to Authenticator App)
    - MFA QR security/backup (for example print)
    - Alternatives 
    - Comparison to other solutions
    - Development Docs???
- Security notes about "always allow" keyring
- Linux & Windows support is essential
- Document auto refresh with supporting aws tools (improvement over broamski)
- Document advisory & mandatory refresh (that matches botocore)
- Document botocore retry (if less than 15*60s expiration)
- Vagrant testing / debugging for Linux & Windows
- Feature comparison chart?
- Docs to https://pkg.go.dev/
- https://medium.com/@jdxcode/12-factor-cli-apps-dd3c227a0e46
- https://github.com/uber-go/guide/blob/master/style.md
- Error handler (At the end)