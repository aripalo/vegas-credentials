# `aws-mfa-credential-process`

🚧 **Work-in-Progress: Do not use just yet! The API and configurations may change without any prior notice at any version. The status of this tool is that it's under development & testing. So do not use this for anything important, but feel free to test this out and give feedback!**

Utility for [AWS `credential_process`](https://docs.aws.amazon.com/sdkref/latest/guide/setting-global-credential_process.html) to assume AWS IAM Roles with _Yubikey Touch and Authenticator App_ [TOPT MFA](https://en.wikipedia.org/wiki/Time-based_One-Time_Password) to provide temporary session credentials – with local caching to [Keyring](#keyring).

If you're unfamiliar with `credential_process`, [this AWS re:Invent video](https://www.youtube.com/watch?v=W8IyScUGuGI&t=1260s) explains it very well.

<br/>

![diagram](/docs/diagram.svg)

<br/>


## Features

- **Fast**: Built with [Go](https://golang.org)

- **Supports automatic temporary session credentials refreshing** for tools that understand session credential expiration

- **Caching of session credentials into [Keyring](#keyring)** to speed things up & to avoid having to input MFA token code for each operation

- **Works out-of-the-box with most AWS tools** such as AWS CDK, AWS SDKs and AWS CLI:

    Tested with AWS CDK (TypeScript), AWS CLI v2, AWS NodeJS SDK (v3), AWS Boto3 and AWS Go SDK. Should probably work with other AWS SDKs as well.

- **Supports both Yubikey Touch or Authenticator App TOPT MFA simultaneously**: 
    
    For example you can default to using Yubikey, but if don't have the Yubikey with you all the time and also have MFA codes in an Authenticator App (such as [Authy](https://authy.com/) for example) you may just type the token code via CLI; Which ever input is given first will be used.

- **Just tap your physical key**:

    No need to manually type or copy-paste TOPT MFA token code from Yubikey Authenticator.

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

    3. more installation methods to be implemented later (maybe pypi?)

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
    assume_role_arn=<target-role-arn>
    credential_process = aws-mfa-credential-process --profile=my-profile
    source_profile=<source-profile-name>
    mfa_serial=<mfa-device-arn>
    yubikey_serial=<yubikey-serial>
    yubikey_label=<yubikey-label>
    ```

5. Use any AWS tooling that support ini-based configuration with `credential_process`, like AWS CLI v2:
    ```shell
    aws sts get-caller-identity --profile my-profile
    ```

<br/>

## Configuration

Configuration for this tool happens `~/.aws/config` ini-file, mostly in the standard way, but do not configure `role_arn`, instead provide `assume_role_arn`: Otherwise AWS tools would ignore the `credential_process` and assume the role directly without using this tool.


### Required default AWS configuration

|       Option        |                                                                                                                                     Description                                                                                                                                      |
| :------------------ | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `source_profile`  | **Required:** Which credentials (profile) are to be used as a source for assuming the target role                                                                                                                                                                                      |
| `mfa_serial`      | **Required:**                                                                                                                                                                                               


### Custom configuration

|       Option        |                                                                                                                                     Description                                                                                                                                      |
| :------------------ | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `assume_role_arn`        | **Required:** The target IAM Role ARN to be assumed.                                                                                                                                                                                                                                         |
| `yubikey_serial`         | **Required if using Yubikey:** Yubikey Device Serial to use. You can see the serial(s) with `ykman list` command. This enforces the use of a specific Yubikey and also enables the support for using multiple Yubikeys (for different profiles)! |
| `yubikey_label`         | **Required if using Yubikey:** Yubikey `oath` Account Label to use. You can see the available accounts with `ykman oath accounts list` command. Set the account label which you have configured your AWS TOPT MFA! |

### Optional default AWS configuration

You can see all the possible configuration options in [AWS documentation](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html#cli-configure-files-global), but mainly you should be interested in:
- `region`
- `duration_seconds`
- `role_session_name`
- `external_id`


### CLI configuratiopn

TODO

--profile
--hide-arns
--verbose


<br/>

## Keyring

In the background this tool uses [`99designs/keyring`](https://github.com/99designs/keyring) to integrate with:
- macOS/OSX Keychain
- Windows credential store
- Pass
- Secret Service
- KDE Wallet
- Encrypted File

🚧 **TODO**: Document keyring usage

<br/>

## TODO


- Configuration file???
- Ensure Role Chaining Works!
- Test manually CDK, CLI, NodeJS SDK v3, Boto3, Go ... for refresh/cache support!
- Add Unit tests
- Add disclaimer for orgs using this tool ("software provided as is")
- Add links to alternatives
- Add PR templates (bug, feature request...)
- Document how to setup Yubikey for TOPT MFA (and additionally add to Authenticator App)
- Document MFA QR security/backup (for example print)
- Blog post
- Development docs (maybe separate site?)
- Contribution guidelines
- Add video that showcases the features (with CDK)
- Comments to code
- TODO log file
- Document credential refresh & botocore multiple calls:
    - https://github.com/aws/aws-cli/issues/5048#issuecomment-597868383
    - https://github.com/boto/botocore/blob/221ffa67a567df99ee78d7ae308c0e12d7eeeea7/botocore/credentials.py#L350-L355