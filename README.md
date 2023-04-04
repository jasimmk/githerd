# Githerd

An application for managing multiple git-repos at once. Multiple inspirations are from [git-repo](https://git-repo.info/en/docs/multi-repos/manifest-format/) project

## Installation


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
