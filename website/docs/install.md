---
sidebar_position: 2
---

# Install

:::info Installation Methods
Additional installation methods may be added later but the below options should cover pretty much every environment.
:::

<br/>

**Via [Homebrew](https://docs.brew.sh/Installation)** on MacOS, GNU/Linux and Windows Subsystem for Linux (WSL):

```sh
brew install aripalo/tap/vegas-credentials
```

**Via [Scoop](https://scoop.sh/)** on Windows:

```sh
scoop bucket add aripalo https://github.com/aripalo/scoops.git && scoop install vegas-credentials
```



## Direct Download

:::caution Software needs to be kept up-to-date
With direct download you lose the possibility of easily [upgrading the software version](#upgrading) and instead you'll have to manually check for newer versions & download them.
:::

In case none of the above installation method work for you, you may go to project's [Github releases](https://github.com/aripalo/vegas-credentials/releases) and download the binary for your platform and add it into your `$PATH`.

## Upgrading

Via Homebrew:
```sh
brew upgrade vegas-credentials
```

Via Scoop:
```sh
scoop update vegas-credentials
```
