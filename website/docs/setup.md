---
sidebar_position: 3
---

# Setup

In this document we're only setting up the necessary configuration changes into `~/.aws/config` to get started using `vegas-credentials` as a credential source.

As in most use cases you do not invoke `vegas-credentials` commands yourself but instead configure it as a credential source for [`credential_process`](https://docs.aws.amazon.com/sdkref/latest/guide/feature-process-credentials.html) in AWS configuration at `~/.aws/config`. If you're interested in the different commands & configuration options passed directly to `vegas-credentials` command, see [manual](/manual).



## Source Profile

1. You must have an AWS IAM User (source) profile credentials configured in `~/.aws/credentials`, for example:
    ```ini
    # ~/.aws/credentials
    [default]
    aws_access_key_id = AKIAIOSFODNN7EXAMPLE
    aws_secret_access_key = wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
    ```

2. As `vegas-credentials` is only meant for assuming roles with MFA, you must configure `mfa_serial` into `~/.aws/config` for your source profile:
    ```ini
    # ~/.aws/config
    [default]
    mfa_serial = arn:aws:iam::111111111111:mfa/FrankSinatra
    ```

## Assumable Profile

Once you've configured your [source profile](#source-profile), you can start defining the assumable profile (target role) in `~/.aws/config`. This is where things get interesting and slighly differentiate from normal AWS configuration:

1. Configure `credential_process` pointing into `vegas-credentials assume` command (with profile name matching the ini section title) with target IAM Role ARN & Source Profile:
    ```ini
    # ~/.aws/config
    [profile frank@concerts]
    credential_process = vegas-credentials assume --profile=frank@concerts
    vegas_role_arn=arn:aws:iam::222222222222:role/SingerRole
    vegas_source_profile=default
    ```

    Note: `role_arn` & `source_profile` MUST be prefixed with `vegas_` to prevent AWS tooling to ignore `credential_process` setting and to prevent Terraform failing.


2. Additionally you may provide any other default AWS configuration options:
    ```ini
    # ~/.aws/config
    [profile frank@concerts]
    credential_process = vegas-credentials assume --profile=frank@concerts
    vegas_role_arn=arn:aws:iam::222222222222:role/SingerRole
    vegas_source_profile=default

    # You may also provide any other additional/optional standard AWS configuration, such as:
    region = us-west-1
    duration_seconds = 4383
    role_session_name = SinatraAtTheSands
    external_id = 0093624694724
    ```


## Testing it out

Once source & assumable profiles are configured, you may use any AWS tools (such as the AWS CLI) to test it out:
```shell
aws sts get-caller-identity --profile frank@concerts
```

Running this command should prompt you for the MFA code which you can input either via GUI Prompt or by typing the code into your Terminal (Standard Input) and pressing enter key.

If the above command succeeds, you should see an output something like below:
```json
{
    "UserId": "AROAUYE2CI3XK7EXAMPLE:SinatraAtTheSands",
    "Account": "222222222222",
    "Arn": "arn:aws:sts::222222222222:role/SingerRole/SinatraAtTheSands"
}
```
