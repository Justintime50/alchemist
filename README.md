<div align="center">

# Alchemist

Update, backup, and administer your Homebrew or Chocolatey instance.

[![Build Status](https://github.com/Justintime50/alchemist/workflows/build/badge.svg)](https://github.com/Justintime50/alchemist/actions)
[![Coverage Status](https://coveralls.io/repos/github/Justintime50/alchemist/badge.svg?branch=main)](https://coveralls.io/github/Justintime50/alchemist?branch=main)
[![Release](https://img.shields.io/github/v/release/Justintime50/alchemist)](https://github.com/Justintime50/alchemist/releases)
[![Licence](https://img.shields.io/github/license/justintime50/alchemist)](LICENSE)

<img src="https://raw.githubusercontent.com/justintime50/assets/main/src/alchemist/showcase.png" alt="Showcase">

</div>

> Bubble bubble, toil and brew...

## Alchemist Backup

Alchemist can backup your entire Homebrew (macOS and Linux) or Chocolatey (Windows) instance. It does this by retrieving the list of installed packages and creating a script that can be run to restore your entire Homebrew or Chocolatey instance.

```bash
alchemist --backup
```

If you run into troubles backing up your Homebrew instance, it's recommended to try running Alcehmist with the `--update` flag first.

## Alchemist Update

**macOS and Linux**

Alchemist automates the entire Homebrew update process including:

1. Updating available taps and formula references
1. Upgrading packages
1. Upgrading casks
1. Cleaning up old/stale taps and formula
1. Checking for problems with your Homebrew instance

**Windows**

Alchemist will update all of your Chocolatey packages:

```bash
alchemist --update
```

## Install

**macOS and Linux**

```bash
# Setup the tap
brew tap justintime50/formulas

# Install the tool
brew install alchemist
```

**Windows**

Download the Windows binary from the [https://github.com/Justintime50/alchemist/releases](releases page).

## Usage

**Logs**

Alchemist saves logs to `~/alchemist/update/alchemist-update.log`. Logs by default are kept on the system for `90 days` and are automatically rotated for you once their size exceeds 1mb or the logs become older than 90 days.

**Restore Scripts**

Scripts generated from the backup functionality of Alchemist live at `~/alchemist/backup`. Simply run `brew bundle --file path/to/Brewfile` or `path/to/restore-choco-packages.bat` to restore your packages.

```
Usage:
    alchemist --update

Options:
    --backup
        Backup your Homebrew instance.
    --update
        Update your Homebrew instance.
    --force
        Forces actions such as backing up even when there are errors. (Brew only)
    --greedy
        Force updates to casks that have auto-update capabilities in their respective UIs.
```

## Development

```bash
# Build the project
make build

# Install the project globally from source
make install

# Clean the executables
make clean

# Test the project
make test

## Get test coverage
make coverage

# Lint the project (requires golangci-lint be installed)
make lint
```
