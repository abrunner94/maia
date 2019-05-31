# Maia

[![Build Status](https://travis-ci.com/anamake/maia.svg?branch=master)](https://travis-ci.com/anamake/maia) [![Go Report Card](https://goreportcard.com/badge/github.com/anamake/maia)](https://goreportcard.com/report/github.com/anamake/maia)

Maia is a CLI and management tool inspired by [pssh](https://pypi.org/project/pssh/), which allows you to execute commands remotely on multiple machines. This is particularly useful if you manage a fleet of machines on-prem and in the cloud, but don't necessarily want to install additional dependencies.

## Get Started

If you would like to compile from source and generate the binaries yourself, first, clone the repository.

```bash
$ git clone https://github.com/anamake/maia.git
```

#### Build and run Maia

```bash
$ cd maia
$ make
```

Before running Maia, make sure to modify your `config.json` to point to the specific machines. Afterwards, run Maia:

```bash
$ maia
```

#### Generating specific binaries

You can also use make to generate a binary for your specific operating system:

#### Linux

```bash
$ make build-linux
```

#### macOS

```bash
$ make build-mac
```

#### Windows

```bash
$ make build-win
```

There are also precompiled binaries available in the Releases section.

#### Why Maia?

Maia was developed due to the constant hassle of checking on particular software versions, machine updates and for executing commands on multiple machines. This is typically done manually by engineers, all of which takes up a lot of time. When adding up the mere lost minutes of tunneling into each machine just to check for a particular software version or execute a specific command, it's time poorly spent.

#### How does it work?

You can configure Maia to work with your machines by defining a `config.json` in the project's root directory. Maia will read the host names, user names, passwords, as well as any SSH keys from this configuration file and executes your user-defined commands remotely on these machines for you.

#### Example

```json
[
  {
    "host": "10.0.0.1",
    "port": "22",
    "user": "user",
    "password": "ssh-password",
    "key": "",
    "command": ["ls", "cat /etc/environment"]
  },
  {
    "host": "10.0.0.2",
    "port": "22",
    "user": "admin",
    "password": "",
    "key": "/path/to/sshkey",
    "command": ["echo $(cat /etc/environment | grep local | cut -d ':' -f1)"]
  },
  {
    "host": "10.0.0.3",
    "port": "22",
    "user": "test",
    "password": "test",
    "key": "",
    "command": ["netstat"]
  }
]
```
