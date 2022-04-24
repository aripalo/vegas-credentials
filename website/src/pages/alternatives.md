# Alternatives

There are many great existing solutions out there that solve similar problems and I've tried to learn from them as much as I can. This tool that I've built is definitely not better or more feature-rich than for example [`99designs/aws-vault`](https://github.com/99designs/aws-vault) in many scenarios as it has a lot more features, more contributors and been around some time. Instead `vegas-credentials` aims to "one thing well": See [Design Principles](/design-principles).

The comparison below focuses on the specific use case this tool tries to solve (i.e. providing a nice UX for assuming a role with MFA using `credential_process` to support as many AWS tools as possible without having to use wrapper scripts).


|                   Feature/Info                   |                                                                                        `aripalo/vegas-credentials`                                                                                         |                                                                [`99designs/aws-vault`](https://github.com/99designs/aws-vault)                                                                 |                                                                 [`broamski/aws-mfa`](https://github.com/broamski/aws-mfa)                                                                 |                                                                [`meeuw/aws-credential-process`](https://github.com/meeuw/aws-credential-process)                                                                 |
| :----------------------------------------------- | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------: | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------: | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------: | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------: |
| Github Stats                                     | ![GitHub Repo stars](https://img.shields.io/github/stars/aripalo/vegas-credentials?style=flat) <br/> ![GitHub last commit](https://img.shields.io/github/last-commit/aripalo/vegas-credentials?style=flat) | ![GitHub Repo stars](https://img.shields.io/github/stars/99designs/aws-vault?style=flat) <br/> ![GitHub last commit](https://img.shields.io/github/last-commit/99designs/aws-vault?style=flat) | ![GitHub Repo stars](https://img.shields.io/github/stars/broamski/aws-mfa?style=flat)  <br/> ![GitHub last commit](https://img.shields.io/github/last-commit/broamski/aws-mfa?style=flat) | ![GitHub Repo stars](https://img.shields.io/github/stars/meeuw/aws-credential-process?style=flat) <br/> ![GitHub last commit](https://img.shields.io/github/last-commit/meeuw/aws-credential-process?style=flat) |
| `credential_process` <br/>with MFA + Assume Role |                                                                                                     ✅                                                                                                      |                                                                                  ❌ [<sup>[*2]</sup>](#note2)                                                                                   |                                                                                ❌ [<sup>[*4]</sup>](#note4)                                                                                |                                                                                                        ✅                                                                                                         |
| Automatic Temporary Session Credential Refresh   |                                                                                                     ✅                                                                                                      |                                                                                  ❌ [<sup>[*3]</sup>](#note3)                                                                                   |                                                                                ❌ [<sup>[*5]</sup>](#note5)                                                                                |                                                                                                        ✅                                                                                                         |
| Yubikey                                          |                                                                                       ✅ ✅ [<sup>[*1]</sup>](#note1)                                                                                        |                                                                                  ✅ [<sup>[*1]</sup>](#note1)                                                                                   |                                                                               ❌  [<sup>[*6]</sup>](#note6)                                                                                |                                                                                          ✅ [<sup>[*10]</sup>](#note10)                                                                                           |
| Cache Encryption                                 |                                                                                                     ✅                                                                                                      |                                                                                               ✅                                                                                                |                                                                               ❌  [<sup>[*7]</sup>](#note7)                                                                                |                                                                                                        ✅                                                                                                         |
| Cache Invalidation on config change              |                                                                                                     ✅                                                                                                      |                                                                                              ✅  ?                                                                                              |                                                                               ✅  [<sup>[*8]</sup>](#note8)                                                                                |                                                                                                        ✅                                                                                                         |
| Cached Performance                               |                                                                                ⚡️ <br/>`<100ms`[<sup>[*11]</sup>](#note11)                                                                                 |                                                                                        ⚡️ <br/>`<50ms`                                                                                         |                                                                            ⚡️ <br/> [<sup>[*9]</sup>](#note9)                                                                             |                                                                                    🐢<br/>`>400ms`[<sup>[*11]</sup>](#note11)                                                                                     |
| Comprehensively Unit Tested                      |                                                                                                     ✅                                                                                                      |                                                                                               ?                                                                                                |                                                                                             ❌                                                                                             |                                                                                                        ✅                                                                                                         |
| Installation methods                             |                                                                                           `brew`, `scoop`                                                                                            |                                                         `brew`, `port`, `choco`, `scoop`, `pacman`, `pkg`, `zypper`, `nix-env`, `asdf`                                                         |                                                                                           `pip`                                                                                           |                                                                                                  `brew`, `pip`                                                                                                   |

Please, [correct me if I'm wrong](https://github.com/aripalo/vegas-credentials/discussions) above or there's any other good alternatives!


## `99designs/aws-vault`

<br/>

1. <a id="note1"></a>Yubikey support in <code>99designs/aws-vault</code> is not perfect:

    - Using multiple Yubikeys is cumbersome due to having to pass in Yubikey [device serial as environment variable for each command](https://github.com/99designs/aws-vault/pull/748) – vs. this tool allows setting device serial via configuration per profile (no need to remember the serial for each Yubikey).
    - Uses deprecated `ykman` commands.
    - See also [point 2](#note2) about `credential_process`, assumed roles and Yubikeys.

<br/>

2. <a id="note2"></a>Does not seem to play well with <code>credential_process</code>:

    - **At least I haven't figured out how to succesfully configure it to use `credential_process`, assume a role, use Yubikey for MFA and to provide temporary session credentials.**
    - They themselves [claim that _“`credential_process` is designed for retrieving master credentials”_](https://github.com/99designs/aws-vault/blob/master/USAGE.md#using-credential_process) - which is NOT true since this tool does work with temporary credentials via `credential_process` just fine and even the [AWS docs on `credential_process` show `SessionToken` and `Expiration` on the expected output from the credentials program](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html).
    - There's further indication that [`99designs/aws-vault` is not designed for `credential_process`](https://github.com/99designs/aws-vault/issues/641#issuecomment-681346113):

        > _**Using credentials_process isn't the way I use aws-vault**, it was a contributed addition, so feels like we should emphasise this is not the recommended path._
        >
        > – Michael Tibben, VP Technology, 99designs

<br/>

3. <a id="note3"></a>This pretty much relates to <a href="#note1">point 1</a>: For AWS tools to automatically request refreshed credentials, the credentials need to be provided via either the multiple standard methods or via <code>credential_process</code>.

<br/>

## `broamski/aws-mfa`

<br/>

4. <a id="note4"></a>Works differently by writing temporary session credentials into <code>~/.aws/credentials</code>, so therefore no <code>credential_process</code> support at all.

<br/>

5. <a id="note5"></a>If temporary session credentials written into <code>~/.aws/credentials</code> by <code>broamski/aws-mfa</code> are expired, AWS tools will fail and you must invoke <code>aws-mfa</code> command manually to fetch new session credentials. There is no (automatic) way for AWS tools to trigger <code>aws-mfa</code> command.

<br/>

6. <a id="note6"></a>You may use Yubikey, but it requires you to manually copy-paste the value from <code>ykman</code> or Yubikey Manager GUI. No "touch integration".

<br/>

7. <a id="note7"></a>Temporary session credentials are written in plaintext into <code>~/aws/credentials</code>. Besides being available as plaintext, it pollutes the credentials file.

<br/>

8. <a id="note8"></a>Configuration is only provided via flags to <code>aws-mfa</code> CLI command, so each time you execute <code>aws-mfa</code> it will use the flags provided. But, the gotcha is that again you need to execute <code>aws-mfa</code> manually always.

<br/>

9. <a id="note9"></a>As temporary session credentials (or "short-term" as <code>aws-mfa</code> calls them) are stored as plaintext into <code>~/aws/credentials</code>, there is no delay since AWS tools can directly read them from that file.

<br/>

## `meeuw/aws-credential-process`

<br/>

10. <a id="note10"></a>Does not support multiple Yubikey devices.

<br/>

11. <a id="note11"></a>Performance

    [Hyperfine](https://github.com/sharkdp/hyperfine) benchmark for retrieving cached temporary session credentials:

    TODO: update

    <!-- ![perf](/docs/perf-comparison.png) -->
