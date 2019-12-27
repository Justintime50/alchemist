# Brew Update

[![Build Status](https://travis-ci.org/Justintime50/brew-update.svg?branch=master)](https://travis-ci.org/Justintime50/brew-update)
[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)

A simple bash script to update Homebrew on macOS.

## Usage

This script can either be run manually or via a cron:

**Manually**
```bash
./update.sh
```

**Cron**
```bash
crontab -e

0 1 * * * /path/to/update.sh
```

*If using a Cron, note that some brew updates cannot be automated and may hang waiting for user input.*
