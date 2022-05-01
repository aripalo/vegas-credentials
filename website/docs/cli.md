---
sidebar_position: 5
---

# CLI

As stated previously, in most cases you don't use the `vegas-credentials` command directly but if you need to troubleshoot something or you wish to configure the user output (such as colors or emoji), then this section is for you.

## Global Flags

|     Flag     | Default Value |                       Purpose                       |
| :----------- | :------------ | :-------------------------------------------------- |
| `--help`     |               | Prints out help message for given command           |
| `--no-color` | `false`       | Disable both colors and emoji from visible output   |
| `--no-emoji` | `false`       | Disable emoji from visible output (but keep colors) |
| `--no-gui`   | `false`       | Disable GUI Diaglog Prompt                          |
| `--verbose`  | `false`       | Enable Verbose Output                               |

## Environment Variables

|             Flag             |                                                                        Purpose                                                                         |
| :--------------------------- | :----------------------------------------------------------------------------------------------------------------------------------------------------- |
| `XDG_CACHE_HOME`             | Enable cache location override as per [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html) |
| `VERBOSE`                    | Enable Verbose Output                                                                                                                                  |
| `TERM=dumb`                  | Disable both colors and emoji from visible output                                                                                                      |
| `NO_COLOR`                   | Disable both colors and emoji from visible output                                                                                                      |
| `VEGAS_CREDENTIALS_NO_COLOR` | Disable both colors and emoji from visible output – only for this program                                                                              |
| `NO_EMOJI`                   | Disable emoji from visible output (but keep colors)                                                                                                    |
| `VEGAS_CREDENTIALS_NO_EMOJI` | Disable emoji from visible output (but keep colors) – only for this program                                                                            |

## Commands

### Root

Mainly for information purposes.

|    Flag     | Required? |                Purpose                 |
| :---------- | :-------: | :------------------------------------- |
| `--version` |           | Prints out version information |
| `--info` |           | Prints out information about the program  |

#### Examples

```sh
$ vegas-credentials --version
v1.21.7
```

```sh
$ vegas-credentials --info
Name:          Vegas Credentials
Version:       v1.21.7
Build:         #abcd123
URL:           https://credentials.vegas
Source:        https://github.com/aripalo/vegas-credentials
Description:   AWS credential_process utility
```

### `assume`

The "main" command, which is used when assigning `vegas-credentials` as the credential source in AWS configuration.

Most of the configuration for this command is read from `~/.aws/config` by looking at the given `--profile` flag.

#### Flags

| Flag | Required? | Purpose |
| :--- | :-----------: | :------ |
|`--profile`| ✓ | Which _Assumable Profile_ to use from `~/.aws/config`. <br/><br/>Should match the value of section title under which `vegas-credentials` is defined as the `credential_process` in `~/.aws/config`.  |

#### Examples

```sh
$ vegas-credentials --profile frank@concerts
{
    "Version": 1,
    "AccessKeyId": "ASIAUYE1FI4W7EXAMPLE",
    "SecretAccessKey": "jvH8V8UQQxLdp0TdacBsfKJoVoYgCGEXAMPLEKEY",
    "SessionToken": "EXAMPLE/SESSION/TOKEN/vYXdzELr//////////wEa89YYBwBWM0EdOBJ2ICKrAVF9fJJcKBk5ez8uzMFUCbUH02FTmq/XvlDPPpBXB/G6Yy7SyAhwFSRyFskurP1aGVdjC/jF3WS1sBVs4r4vf5udPC/kJJiox/a+xk4Z0ZfXy139vtfbdrBjw1mSVNzhW/gXcZbRRhKJMEl+7vDGNiQ0MqZa1Fz0E26s40av4F2BQac0jnOSqE8GazgCeRjUyxgHtwHivEEKwiQDxjj5W7f9AM56RSyJlByj/3JCTBjItfwFv3qEJb5cu1pe/r1RnRNVzdgqGbc+Y1Mr1x+EXAMPLE",
    "Expiration": "2022-04-21T17:51:11Z"
}
```

### `config list`

🚧 **Not implemented yet!**

Prints out information what are the various configuration, cache and temporary file locations in use.


### `cache clean`

🚧 **Not implemented yet!**

Cleans up the cached data.

| Flag | Required? | Purpose |
| :--- | :-----------: | :------ |
|`--password`|  | Deletes the Yubikey OATH application password cache only  |
|`--credential`|  | Deletes the Temporary Session Credential cache only  |
