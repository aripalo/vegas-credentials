---
sidebar_position: 19
---

# Examples

## Multiple Source Profiles

Let's say you have an AWS IAM user for work stuff and another for your hobby project, both which you use to assume roles. This can be achieved with following example:
```ini
# ~/.aws/config
[profile work]
mfa_serial = arn:aws:iam::111111111111:mfa/FrankSinatra

[profile hobby]
mfa_serial = arn:aws:iam::999999999999:mfa/Frankie

[profile frank@concerts]
credential_process = vegas-credentials assume --profile=frank@concerts
vegas_role_arn=arn:aws:iam::222222222222:role/SingerRole
vegas_source_profile=work

[profile frankie@painting]
credential_process = vegas-credentials assume --profile=frankie@painting
vegas_role_arn=arn:aws:iam::888888888888:role/PainterRole
vegas_source_profile=hobby
```


## AWS SDKs

Often times you may not want to define the `profile` within the application code, since the application code most often will be ran without profile in cloud. You may circumvent this by setting the profile via `AWS_PROFILE` environment variable. Most AWS SDKs should support this. Example with running a NodeJS based script (that does something with AWS SDK):
```shell
AWS_PROFILE=frank@concerts node index.js
```

> _By default, the SDK checks the `AWS_PROFILE` environment variable to determine which profile to use. If the `AWS_PROFILE` variable is not set in your environment, the SDK uses the credentials for the `[default]` profile. To use one of the alternate profiles, set or change the value of the `AWS_PROFILE` environment variable. For example, given the configuration file shown, to use the credentials from the work account, set the `AWS_PROFILE` environment variable to work-account (as appropriate for your operating system)._
>
> – [AWS SDK v3 docs](https://docs.aws.amazon.com/sdk-for-javascript/v3/developer-guide/loading-node-credentials-shared.html)

## Role Chaining

This tool also supports role chaining - **given that the specific AWS tool your using supports it** - which means assuming an initial role and then using it to assume another role. An example with 3 different AWS accounts would look like:

![role-chaining](/img/role-chaining.svg)

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

## Terraform

Use Terraform AWS Provider to define your resources and define variable `aws_profile`:
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

variable "aws_profile" {
  type = string
}

# Again, nothing special here, just normal profile configuration…
provider "aws" {
  profile = var.aws_profile
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

Run Terraform and define the profile (for example via CLI `-var` flag):
```bash
terraform apply -var="aws_profile=frank@concerts
```

### Parallelism

Terraform [performs operations in parallel](https://www.terraform.io/cli/commands/apply#parallelism-n) which means that it ends up invoking the `credential_process` (and therefore `vegas-credentials assume` command) multiple times concurrently; This is okay, as this tool now has support for parallelism where each process invocation uses [`go-filemutex`](https://github.com/alexflint/go-filemutex) to obtain a file lock or wait until a lock can be acquired (i.e. wait until another process invocation has finished): This results into a queued authentication flow where the first invocation of `vegas-credentials assume` prompting for MFA token (unless valid temporary session credentials were already available from cache) and consecutive invocations receiving the temporary session credentials from cache.

## Ansible

Define your AWS resources with Ansible as usual, but add `profile` key in which you may use a variable that could be set via command-line `--extra-vars` for example:

```yaml
---
- name: Create Bucket
  hosts: local
  connection: local
  vars:
    aws_profile:
  tasks:
    - name: Create new bucket
      aws_s3:
        bucket: franks-bcuket
        mode: create
        region: us-west-1
        profile: "{{ aws_profile }}"
```

Run it:
```bash
ansible-playbook -i hosts tasks.yml --extra-vars "aws_profile=frank@concerts"
```
