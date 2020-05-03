<div align="center">

# Brew Update

Automate the entire Homebrew update process with just one command.

[![Build Status](https://travis-ci.com/Justintime50/brew-update.svg?branch=master)](https://travis-ci.com/Justintime50/brew-update)
[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)

<img src="assets/showcase.gif">

</div>

## Usage

Brew Update automates the entire Homebrew update process including updating available taps and formulas, upgrading packages, upgrading cask packages, cleaning up old/stale taps and formulas, and checking for problems with your Homebrew instance. The script saves the output to a log file on each run found at `~/brew-update` (this can be customized).

If you setup Brew Update as a Launch Agent, it will run whenever it's loaded (on login). This is a great method to use if you want to have your brew setup updated at the beginning of your usage so you are always up to date.

There's lots of ways to use Brew Update, check them out below:

### Run Script

```bash
./brew-update.sh
```

### Shell Alias

```bash
# If using Bash insted of ZSH, use ~/.bash_profile
echo alias brew-update="/path/to/brew-update.sh" >> ~/.zshrc
source ~/.zshrc

# Usage of alias
brew-update
```

### Launch Agent (Recommended on macOS for automation)

Edit the path in the `plist` file to your script and logs as well as the time to execute, then setup the Launch Agent:

```bash
# Copy the plist to the Launch Agent directory
cp local.brewUpdate.plist ~/Library/LaunchAgents

# Use `load/unload` to add/remove the script as a Launch Agent
launchctl load ~/Library/LaunchAgents/local.brewUpdate.plist

# To `start/stop` the script from running, use the following
launchctl start local.brewUpdate.plist
```

### Cron

```bash
crontab -e

0 1 * * * /path/to/brew-update.sh
```
