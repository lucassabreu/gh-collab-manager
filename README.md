gh-collab-manager
=================

A cli to help invite and/or remove multiple collaborators from multiple GitHub repositories.

[![Release](https://img.shields.io/github/release/lucassabreu/gh-collab-manager.svg?classes=badges)](https://github.com/lucassabreu/gh-collab-manager/releases/latest)
[![Build Status](https://github.com/lucassabreu/gh-collab-manager/actions/workflows/release.yml/badge.svg?classes=badges)](.github/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/lucassabreu/gh-collab-manager?classes=badges)](https://goreportcard.com/report/github.com/lucassabreu/gh-collab-manager)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Flucassabreu%2Fgh-collab-manager.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Flucassabreu%2Fgh-collab-manager?ref=badge_shield)

Usage
-----

Invite user `johndue` to `lucassabreu/gh-collab-manager` and `lucassabreu/clockify-cli`.

```sh
gh-collab-manager --github-token <token> -a johndue \
  -R lucassabreu/gh-collab-manager -R lucassabreu/clockify-cli
```

Remove `johndue` from `lucassabreu/gh-collab-manager`.

```sh
gh-collab-manager --github-token <token> -R lucassabreu/gh-collab-manager -r johndue
```

Invite `octocat` and `github`; also remove `johndue` from repositories `lucassabreu/gh-collab-manager` and `lucassabreu/clockify-cli`.

```sh
gh-collab-manager --github-token <token> \
  -R lucassabreu/gh-collab-manager -R lucassabreu/clockify-cli \
  -r johndue -a octocat -a github
```

How to install [![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?classes=badges)](https://github.com/goreleaser)
--------------

#### Using `go install`

```sh
go install github.com/lucassabreu/gh-collab-manager
```

#### By Hand

Go to the [releases page](https://github.com/lucassabreu/gh-collab-manager/releases) and download the pre-compiled
binary that fits your system.

Help
----

```console
Add and remove users from repositories

Usage:
  gh-collab-manager [flags]

Flags:
  -a, --add stringArray       handle of collaborators to invite to the repositories
      --config string         config file (default is $HOME/.gh-collab-manager.yaml)
  -t, --github-token string   github token to access the api (defaults to $GITHUB_TOKEN)
  -h, --help                  help for gh-collab-manager
  -r, --remove stringArray    handle of collaborators to remove from the repositories
  -R, --repo stringArray      repositories to add/remove collaborators
  -v, --version               shows cli version
```

Changelog
---------

[Link](CHANGELOG.md)


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Flucassabreu%2Fgh-collab-manager.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Flucassabreu%2Fgh-collab-manager?ref=badge_large)