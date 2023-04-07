# Githerd

An application for managing multiple git-repos at once.

## Installation

There are two ways to install githerd, either from source or from binary. Staight forward way is to download the binary from the [releases](https://github.com/jasimmk/githerd/) page.

### From source

```bash
go install github.com/jasimmk/githerd/cmd/githerd
```

## Commands

All the commands are listed in the [documentation page](./docs/commands/.md).



## Set config

Create a file `~/.githerd/config.yaml` with the following contents in case you doesn't want to use PR creation or mirror functionality:

```yaml
---
```

if you want to use PR creation or mirror functionality, you need to add the following to the config file:

```yaml
---
profiles:
  - name: github-profile
    api_type: github # github/bitbucket/gitlab
    ssh_key: <keyfile>
    token: <token>
```

Licenses are listed in the [`LICENSE`](LICENSE) file.
Credits and inspirations are listed in the [`CREDITS`](CREDITS.md) file.