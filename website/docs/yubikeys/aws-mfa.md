---
sidebar_position: 2
---

# AWS MFA Setup

Using Yubikey(s) with `vegas-credentials` in AWS requires you to set up a new _Virtual MFA Device_ in AWS.

**The use of _Virtual MFA Device_ is required since we're dealing with OATH Time-based One-Time Passwords (TOTP) MFA for multifactor authentication via CLI** – and NOT ~~_U2F Device_~~ multifactor authentication which can only be used via the Web Console UI.

1. Go to [AWS IAM Users](https://console.aws.amazon.com/iamv2/home#/users) page.

2. Find your user and choose `Security credentials` tab.

3. In the _Assigned MFA Device_ row, choose `Manage`:

    ![manage mfa](/img/yubikey-setup/1-aws-manage-mfa.png)

4. From the new modal dialog choose _Virtual MFA Device_:

    ![virtual mfa](/img/yubikey-setup/2-aws-add-virtual-mfa.png)

5. For easiest setup, click _Show QR_ and use Yubico Authenticator GUI to setup:

    ![show-qr](/img/yubikey-setup/3-aws-qr-hidden.png)

    Alternatively you may select _Show secret key_ and use `ykman` CLI to set up the MFA. For reference you may follow [these instructions' step 6 and 7](https://aws.amazon.com/blogs/security/enhance-programmatic-access-for-iam-users-using-yubikey-for-multi-factor-authentication/).

6. At this point it's useful to scan the QR also into your mobile Authetincator app (such as [Authy](https://authy.com/) or [Google Authenticator](https://support.google.com/accounts/answer/1066447)) so they act as a backup. Do NOT use the MFA Codes from the mobile app to configure AWS at this point!

7. **IF you're configuring the MFA here for the AWS Account _Root_ itself, then you definitely should strongly think about MFA backup so that you won't lose access into your account!**

    - Besides using a mobile Authenticator App as a backup, consider saving the QR code or the secret key into a secure location! For example you could print the QR code and physically save it in a secure location and/or save it into a some kind of secure secrets management system. These are just examples! The root account MFA backup strategy is an important concepts and it is up to you or your organization to define the best & most secure option!


8. Open Yubico Authenticator and add the new account:

    | 1.  | 2. | 3. |
    |:--:|:--:|:--:|
    | ![1](/img/yubikey-setup/4-yubico-scan-qr.png) | ![2](/img/yubikey-setup/5-yubico-add-account.png) | ![3](/img/yubikey-setup/6-yubico-code-copied.png) |

    1. You may also add the account details manually, but the QR scan is the fastest. The QR code in AWS Console UI must be visible to do this.

    2. Once account is added, the issuer should be `Amazon Web Services` and Account follows the pattern `<YourUserName>@<YourAwsAccountAlias>`. **You MUST leave `Require touch` as enabled!**

9. Paste 2 consecutive distinct MFA codes into the AWS from the Yubico Authenticator:

    ![add codes](/img/yubikey-setup/7-aws-add-codes.png)

10. Now the MFA setup in AWS is done:

    ![mfa-ready](/img/yubikey-setup/8-aws-mfa-ready.png)

    You should sign out and ensure you can sign again into the AWS Web Consule UI with both MFA from the Yubico Authenticator application and from your backup mobile Authenticator App!
