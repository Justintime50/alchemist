<div align="center">

# Brew Update

Automate the entire Homebrew update process with just one command.

[![Build](https://github.com/Justintime50/brew-update/workflows/build/badge.svg)](https://github.com/Justintime50/brew-update/actions)
[![Licence](https://img.shields.io/github/license/justintime50/brew-update)](LICENSE)

<img src="assets/showcase.gif" alt="Showcase">

</div>

Brew Update automates the entire Homebrew update process including:

1. Updating available taps and formulas
1. Upgrading packages
1. Upgrading casks
1. Cleaning up old/stale taps and formulas
1. Checking for problems with your Homebrew instance

The script saves the output on each run to a log file found at `~/brew-update` by default. Logs by default are kept on the system for `90 days` and are cleaned up each time this tool is run.

## Install

```bash
# Setup the tap
brew tap justintime50/formulas

# Install the tool
brew install brew-update
```

## Usage

Update your entire Homebrew instance with the following command. Pass an optional custom path to save logs to as the first parameter and an optional custom number for the log life in days as the second parameter.

```bash
brew-update ~/custom_path 14
```
