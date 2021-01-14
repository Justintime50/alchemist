<div align="center">

# Alchemist

Automate the entire Homebrew update process with just one command.

[![Build](https://github.com/Justintime50/alchemist/workflows/build/badge.svg)](https://github.com/Justintime50/alchemist/actions)
[![Licence](https://img.shields.io/github/license/justintime50/alchemist)](LICENSE)

<img src="assets/showcase.gif" alt="Showcase">

</div>

Brew Update automates the entire Homebrew update process including:

1. Updating available taps and formulas
1. Upgrading packages
1. Upgrading casks
1. Cleaning up old/stale taps and formulas
1. Checking for problems with your Homebrew instance

The script saves the output on each run to a log file found at `~/alchemist` by default. Logs by default are kept on the system for `90 days` and are cleaned up each time this tool is run.

## Install

```bash
# Setup the tap
brew tap justintime50/formulas

# Install the tool
brew install alchemist
```

## Usage

```
Usage:
    alchemist --update

Options:
    -backup
        Backup your Homebrew instance.
    -update
        Update your Homebrew instance.
```
