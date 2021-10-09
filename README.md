# `aws-mfa-assume-credential-process`

A helper utility that plugs into standard [`credential_process`](https://docs.aws.amazon.com/sdkref/latest/guide/setting-global-credential_process.html) to assume AWS IAM Role with _– Yubikey Touch or Authenticator App –_ MFA to provides session credentials.

<br/>

![diagram](/docs/diagram.svg)

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

This `aws-mfa-assume-credential-process` is _yet another tool_, but it plugs into the standard [`credential_process`](https://docs.aws.amazon.com/sdkref/latest/guide/setting-global-credential_process.html) AWS configuration so most of AWS tooling (CLI v2, SDKs and CDK) will work out-of-the-box with it and also support automatic temporary session credential refreshing.


## Getting Started

1. NodeJS `v14` or newer required

2. [Install `ykman`](https://developers.yubico.com/yubikey-manager/) (if you choose to use Yubikeys)

2. Install:

    ```shell
    npm i -g aws-mfa-assume-credential-process
    ```

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
    credential_process = aws-mfa-assume-credential-process --source=<source-profile-name> --assume=<target-role-arn>
    ```

5. Use any AWS tooling that support ini-based configuration with `credential_process`, like AWS CLI v2:
    ```shell
    aws sts get-caller-identity --profile my-profile
    ```

## Configuration

| Command-line Option |                                                                                                                                  Description                                                                                                                                   |
| :------------------ | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `source`            | *Required:* Which credentials (profile) are to be used as a source for assuming the target role                                                                                                                                                                                                      |
| `assume`            | *Required:* The target IAM Role to be assumed                                                                                                                                                                                                                                              |
| `region`            | Which AWS region to use, if not provided it will use your default AWS region                                                                                                                                                                                                   |
| `duration`          | The value can range from `900` seconds (15 minutes) up to the maximum session duration setting for the role (which can be a maximum of `43200`). This is an optional parameter and by default, the value is set to `3600` seconds.                                             |
| `session-name`      | Specifies the name to attach to the role session. By default this tool will generate a session name based on your source credentials                                                                                                                                           |
| `external-id`       | Specifies a unique identifier that is used by third parties to assume a role in their customers' accounts. This maps to the ExternalId parameter in the AssumeRole operation. This parameter is needed only if the trust policy for the role specifies a value for ExternalId. |
| `yubikey`           | Enable Yubikey usage by providing the Yubikey Device Serial to use. You can see the serial(s) with `ykman list` command. This enforces the use of a specific Yubikey and also enables the support for using multiple Yubikeys!                                                 |

### Example

An example with all the configuration options:
```ini
[profile my-profile]
credential_process = aws-mfa-assume-credential-process --source=default --assume=arn:aws:iam::999988887777:role/MyTargetRole --region=eu-west-1 --duration=900 --session-name=mySession --external-id=foobar --yubikey=12345678
```

### Why CLI options and not just use the default ini-configuration?

If we provide (target) `role-arn` and other configuration in the `~/.aws/config` ini-file the standard way, then most AWS tools will ignore the `credential_process` and assume the role directly without using this tool.



## TODO

- Ensure CDK & co understand the session credential expiration and do not ask for MFA all the time
- Document TTY usage https://github.com/boto/botocore/issues/1348
- Support role chaining?
- Add disclaimer for orgs using this tool ("software provided as is")
- Add PR templates (bug, feature request...)