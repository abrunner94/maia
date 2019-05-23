# Maia

Maia is a CLI that allows you to execute remote commands on multiple machines at once. This is particularly useful if you manage a fleet of machines on-prem and in the cloud.

#### Why Maia?

Maia was developed due to the constant hassle of checking on particular software versions, machine updates and remote executions that are typically done manually by engineers. This takes time, and when adding up the mere lost minutes of tunneling into each machine just to check for a particular software version, it's quite phenomenal and time poorly spent.

#### How does it work?

This is primarily done by using a `config.json` file and Go's native SSH client library. Maia will read the host names, user names, passwords, as well as any SSH keys from this configuration file, and executes user-defined commands remotely on these machines.
