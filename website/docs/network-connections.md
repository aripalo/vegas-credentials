---
sidebar_position: 20
---


# Network Connections

If you're using a firewall software (such as [Little Snitch](https://obdev.at/products/littlesnitch/) on macOS) you need to allow following outbound connections to be made by `vegas-credentials` process:

| Type  | Port  |       Target        |                                 Reason                                 |
| :---- | :---- | :------------------ | :--------------------------------------------------------------------- |
| DNS   | `53`  | local network       | DNS resolving                                                          |
| HTTPS | `443` | `sts.amazonaws.com` | Authenticating the IAM User and fetching Temporary Session Credentials |

