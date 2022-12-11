---
sidebar_position: 3
---

# Yubikey(s) with Vegas

1. In your `~/.aws/config` you should have the `mfa_serial` configured and matching the _Virtual MFA Device_ serial we configured in the previous step:
    ```ini
    # ~/.aws/config
    [default]
    mfa_serial = arn:aws:iam::111111111111:mfa/FrankSinatra
    ```

2. Configure the Yubikey OATH account "label":
    ```ini
    # ~/.aws/config
    [default]
    mfa_serial = arn:aws:iam::111111111111:mfa/FrankSinatra
    vegas_yubikey_label = Amazon Web Services:FrankSinatra@vegas-demo-account
    ```

    The format usually follows `<issuer>:<account-name>`. If you added the account via `ykman` CLI it is possible to use other formats as well: One popular choice is to use the same value as `mfa_serial` (such as `arn:aws:iam::111111111111:mfa/FrankSinatra`), in that case you don't need to provide the `vegas_yubikey_label` configuration options and `vegas-credentials` will automatically use the value of `mfa_serial` as the account label.










## Multiple Yubikeys

**IF** you are using multiple Yubikey Devices, you must configure the _Device Serial Number_ into your `~/.aws/config` as follows:
```ini
# ~/.aws/config
[default]
mfa_serial = arn:aws:iam::111111111111:mfa/FrankSinatra
vegas_yubikey_serial = 12345678
```

### Multiple Source Profiles with different Yubikeys

This is especially useful if you have let's say separate work and hobby AWS accounts you use as the source profiles and you also have two different Yubikeys you use for work and hobbies.
```ini
# ~/.aws/config
[profile work]
mfa_serial = arn:aws:iam::111111111111:mfa/FrankSinatra
vegas_yubikey_serial = 12345678

[profile hobby]
mfa_serial = arn:aws:iam::999999999999:mfa/Frankie
vegas_yubikey_serial = 87654321
```


### Multiple Source Profiles for same AWS account but different MFA options

This is especially useful if you want the possibility of using multiple MFA options for single AWS account, like
* Google Authenticator Virtual MFA device
* Yubikey Virtual MFA device

To support this, you need the following
```ini
# ~/.aws/config

#
# Two "credential profiles" for same account and user, but different virtual MFA device 
#

[profile work-authenticator]
mfa_serial = arn:aws:iam::111111111111:mfa/FrankSinatra
# No vegas_yubikey_label as this MFA serial is for Google Authenticator app based MFA

[profile work-yubikey]
mfa_serial = arn:aws:iam::111111111111:mfa/FrankSinatra@yubikey-oath
vegas_yubikey_label = Amazon Web Services:FrankSinatra@vegas-demo-account

#
# Two "role profiles" for assuming a role using the above "credential profiles".
# These profiles assume the same role, but they use different credential profiles so that you have the optin of using either Google Autnehticator or Yubikey
#

[profile somerole@work-acc-2@a]
credential_process = vegas-credentials assume --profile=somerole@work-acc-2@a
vegas_role_arn=arn:aws:iam::2222222222222:role/somerole
vegas_source_profile=work-authenticator

[profile somerole@work-acc-2@y]
credential_process = vegas-credentials assume --profile=somerole@work-acc-2@y
vegas_role_arn=arn:aws:iam::2222222222222:role/somerole
vegas_source_profile=work-yubikey
```

The mfa_serial values above must correspond to the value in IAM -> User -> Security Credentials -> Multi-factor authentication (MFA) list.
The vegas_yubikey_label must match the name you gave it in Yubikey Authenticator GUI and it must follow the pattern Amazon Web Services:<user>@<account-alias>


In addition to setting up these two profiles in ~/.aws/config you also must set up corresponding profiles in ~/.aws/credentials

```ini
# ~/.aws/credentials
[work-authenticator]
aws_access_key_id = AKIATH3ACC3SSK3Y
aws_secret_access_key = th353cr3tk3y+/2353523523

[work-yubikey]
aws_access_key_id = AKIATH3ACC3SSK3Y
aws_secret_access_key = th353cr3tk3y+/2353523523
```

Notice that the aws_access_key_id and aws_secret_access_key values are the same for these two profiles.

After this setup, you are able to perform the following to check the "credential profiles" are configured correctly:
```shell
aws sts get-caller-identity --profile work-authenticator
{
    "UserId": "AIDAS4636346546546",
    "Account": "111111111111",
    "Arn": "arn:aws:iam::111111111111:user/FrankSinatra"
}

aws sts get-caller-identity --profile work-yubikey
{
    "UserId": "AIDAS4636346546546",
    "Account": "111111111111",
    "Arn": "arn:aws:iam::111111111111:user/FrankSinatra"
}

```

And finally, you are now able to use vegas-credentials to assume the same role using either Google Authenticator or Yubikey.

Following asks for Google Authenticator TOTP:
```shell
aws s3 ls --profile somerole@work-acc-2@a
```

Following asks for touching the yubikey to get the TOTP:
```shell
aws s3 ls --profile somerole@work-acc-2@y
```




