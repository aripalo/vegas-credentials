---
sidebar_position: 7
---

# Privacy Policy

**`vegas-credentials` does not use or send any tracking, metrics or analytics data.**

For troubleshooting purposes, some runtime events are written into [`application.log`-file](/docs/logs) on your own machine, but contents of that log file are kept on your machine – unless you yourself explicitly transfer them elsewhere. Also sensitive information is NOT written to that log file.

To enable smooth user experience, Vegas Credentials caches AWS STS Temporary Session Credentials and Yubikey OATH application passwords (if configured). Cached data is:
- persisted in your computer's local filesystem
- always set to expire using Time-To-Live attributes
- encrypted with `AES-256-CTR` (but with system derived secrets)

Read [Design Principles](/design-principles#cache-mechanism) to learn more about the caching mechanism.

Only [network connections](/docs/network-connections) made by Vegas Credentials are the ones needed to authenticate with AWS IAM and request temporary session credentials from AWS STS (using [AWS Go SDK](https://github.com/aws/aws-sdk-go)).

Vegas Credentials is open source software published _“‘AS IS’ BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND”_ with [Apache License 2.0](https://github.com/aripalo/vegas-credentials/blob/main/LICENSE) and the [source code is available on Github](https://github.com/aripalo/vegas-credentials/) for anyone to investigate or to fork. Binary releases are made with [GoReleaser](https://goreleaser.com/) and the release artifacts with their checksums are available on [Github Releases](https://github.com/aripalo/vegas-credentials/releases).

See also [Design Principles](/design-principles).
