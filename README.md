<div align="center">

# Alchemist

Update, backup, and administer your Homebrew instance.

[![Build](https://github.com/Justintime50/alchemist/workflows/build/badge.svg)](https://github.com/Justintime50/alchemist/actions)
[![Coverage Status](https://coveralls.io/repos/github/Justintime50/alchemist/badge.svg?branch=main)](https://coveralls.io/github/Justintime50/alchemist?branch=main)
[![Licence](https://img.shields.io/github/license/justintime50/alchemist)](LICENSE)

<img src="assets/showcase.png" alt="Showcase">

</div>

> Bubble bubble, toil and brew...

## Alchemist Backup

Alchemist can backup your entire Homebrew instance. It does this by retrieving the list of Homebrew packages and casks and creating shell scripts that can be run to restore your entire Homebrew instance.

```bash
alchemist --backup
```

## Alchemist Update

Alchemist automates the entire Homebrew update process including:

1. Updating available taps and formula references
1. Upgrading packages
1. Upgrading casks
1. Cleaning up old/stale taps and formula
1. Checking for problems with your Homebrew instance

```bash
alchemist --update
```

## Install

```bash
# Setup the tap
brew tap justintime50/formulas

# Install the tool
brew install alchemist
```

## Usage

Alchemist saves logs to `~/alchemist`. Logs by default are kept on the system for `90 days` and are automatically rotated for you once their size exceeds 1mb or the logs become older than 90 days.

```
Usage:
    alchemist --update

Options:
    -backup
        Backup your Homebrew instance.
    -update
        Update your Homebrew instance.
```

## Attribution

* Icons made by <a href="https://www.flaticon.com/free-icon/chemist_2646063?term=chemist&related_id=2646063" title="ultimatearm">ultimatearm</a> from <a href="https://www.flaticon.com/" title="Flaticon">www.flaticon.com</a>
