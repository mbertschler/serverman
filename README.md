# `serverman` Server Setup Tooling

Serverman (server manager) is a framework and tooling around setting up software on servers using Go code.

## Why

The goal of `serverman` is to provide a simple way to install software on a single server and keep it up to date.
You declare a configuration of things you want to be installed using Go code.
It chooses a statically typed programming language instead of config files to take advantage of the benefits
of type checking and compilation.
This also allows you to easily add custom logic which is harder to do in a config file based tool.
It also makes it possible to use Go modules for importing a specific version of third party config packages.

## Should I use it?

If you want to

- configure the software you want to install in Go
- manage your config in a Git repository
- automatically keep your installed packages up to date
- use a Debian/Ubuntu based Linux distribution

you should give `serverman` a try.

## How does it work?

Starting from a fresh OS install you run a shell script that downloads the `serverman` bootstrap tool and provide
a Git URL to a repository where your personal config is stored.

- the bootstrap tool sets up a Go toolchain if required
- it checks out your config repository
- it builds and runs your config code
- it regularly checks if new commits can be pulled
- if there are new commits, it rebuilds the config

The tool keeps a state of all applied installs so that it can remove software that is not part of a newer config.
It prefers using CLIs over APIs for ease of debugging, manual testing and keeping the list of dependencies small.

## Planned Config Packages

- [ ] `apt` - install and update Debian packages
- [ ] `systemd` - manage services
- [ ] `download` - download files and unpack them
- [ ] `git` - clone and pull of Git repositories
- [ ] `go` - install using `go get` and modules
- [ ] `nomad` - Hashicorp Nomad client
