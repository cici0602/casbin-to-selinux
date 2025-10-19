# Casbin to SELinux

# test test

[![Go Report Card](https://goreportcard.com/badge/github.com/casbin/casbin-to-selinux)](https://goreportcard.com/report/github.com/casbin/casbin-to-selinux)
[![Build](https://github.com/casbin/casbin-to-selinux/actions/workflows/ci.yml/badge.svg)](https://github.com/casbin/casbin-to-selinux/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/casbin/casbin-to-selinux)](https://github.com/casbin/casbin-to-selinux/releases)
[![Discord](https://img.shields.io/discord/1022748306096537660?logo=discord&label=discord&color=5865F2)](https://discord.gg/S5UjpzGZjN)

Lightweight compiler that translates Casbin PML (Policy Modeling Language) into SELinux policy modules.

## Features

Supported:

- Basic Type Enforcement (TE)
- File and directory access control
- TCP/UDP port bind/connect
- Unix domain sockets
- Process capabilities
- Domain transitions
- Generate standard SELinux files (.te, .fc, .if)

Not supported:

- MLS/MCS (multi-level security)
- Conditional policies and booleans
- Complex constraint statements
- Role transitions

## Supported Systems

This project generates standard SELinux policy files compatible with Linux distributions that use standard SELinux implementation, such as:

- Red Hat Enterprise Linux (RHEL)
- Fedora
- CentOS
- SUSE Linux Enterprise Server (SLES)

It does not support Android SELinux or other non-standard variants.

## Installation

Install the CLI using go:

```bash
go install github.com/casbin/casbin-to-selinux/cli@latest
```

Or build from source:

```bash
git clone https://github.com/casbin/casbin-to-selinux.git
cd casbin-to-selinux
make build
```

## Quick start

1. Initialize a project:

```bash
casbin2selinux init myapp
cd myapp
```

2. Edit `policy.csv` (example):

```csv
# Allow executing the application binary
p, myapp_t, /usr/bin/myapp, execute, allow

# Allow reading the configuration file
p, myapp_t, /etc/myapp/config.json, read, allow

# Allow appending to logs
p, myapp_t, /var/log/myapp(/.*)?, append, allow

# Allow binding TCP port 8080
p, myapp_t, tcp:8080, name_bind, allow
```

3. Compile the policy:

```bash
casbin2selinux compile -m model.conf -p policy.csv -o output/
```

## Commands

```bash
casbin2selinux compile    # Compile PML to SELinux policy
casbin2selinux validate   # Validate PML files
casbin2selinux init       # Initialize a new project
casbin2selinux version    # Show version
```

## Examples

See the `examples/` directory for sample projects:

- `webapp/` - web application example
- `database/` - database service example
- `worker/` - background worker example

## License

Apache License 2.0
