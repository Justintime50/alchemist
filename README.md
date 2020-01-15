<div align="center">

# Brew Update

Automate updating Homebrew with one easy script.

[![Build Status](https://travis-ci.org/Justintime50/brew-update.svg?branch=master)](https://travis-ci.org/Justintime50/brew-update)
[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)

<img src="assets/showcase.gif">

</div>

## Usage

Brew Update automates the entire Homebrew update process including updating available taps and formulas, upgrading packages, upgrading cask packages, cleaning up old/stale taps and formulas, and running brew doctor. It saves the output to a log file on each run found at `~/brew-update` (this can be customized).

This script can either be run manually or via a cron:

### Manually
```bash
./update.sh
```

### Cron
```bash
crontab -e

0 1 * * * /path/to/update.sh
```

*If using a Cron, note that seldom some brew updates cannot be automated and may hang waiting for user input.*
