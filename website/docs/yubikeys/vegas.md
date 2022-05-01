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
