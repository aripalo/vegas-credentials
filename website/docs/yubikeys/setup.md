---
sidebar_position: 1
---

# Yubikey Setup

This section covers how you can configure and use Yubikeys as one of the MFA authentication methods with `vegas-credentials`. The nice thing is the _“one of the methods”_ part: You can (and should) set your MFA to multiple devices, for example into one Yubikey and into your mobile Authenticator App; This way if you don't have your Yubikey device with you at all times, you'd still be able to authenticate with your mobile device!

## Prerequsites

1. You MUST have at least one [Yubikey Touch device](https://www.yubico.com/products/yubikey-5-overview/) with [OATH TOTP](https://en.wikipedia.org/wiki/Time-based_One-Time_Password) support (Yubikey 5 or 5C recommended).

2. You MUST have Yubikey Manager CLI [`ykman` CLI](https://developers.yubico.com/yubikey-manager/) installed in your machine and available in your `$PATH`.

3. You MAY also install the [Yubico Authenticator](https://www.yubico.com/products/yubico-authenticator/) GUI which helps you when you're setting up new OATH accounts into your Yubikey.

4. You SHOULD consider password protecting OATH application on your Yubikey device. You can do that with `ykman`:
    ```sh
    ykman oath access change
    ```

    1. For extra protection, you SHOULD NOT use `ykman oath access remember` or click “Remember password” with the Yubico Authenticator GUI.

    2. For password protected devices, `vegas-credentials` will prompt you for the device OATH application password and cache it for 12 hours: Which usually is enought to get you through the work day without having to retype the password.

