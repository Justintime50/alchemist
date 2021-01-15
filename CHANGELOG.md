# CHANGELOG

## v2.0.0 (TODO)

* Renamed `brew-update` to `alchemist`
* Consolidated `brew-backup` tool into Alchemist, now you can specify the functionality by passing a flag (eg: `--update`, `--backup`)
* Rewrote the tool from shell scripts into Golang
* Removed the reliance on system tools like `sed` for much safer and compatible functionality
* Added proper logging framework with automatic rollover (90 day log life, 1mb file size, 5 logs max)

## v1.2.1 (2021-01-09)

* Switching from Travis-CI to GitHub Actions
* Using `homebrew-releaser` for automated releasing

## v1.2.0 (2020-10-05)

* Refactored code to use functions
* Added parameters to allow for a custom log path and log life in days

## v1.1.0 (2020-09-16)

* Now installable via Homebrew
* Removed launch agent
* Replaced `brew cask upgrade` with `brew upgrade --cask` as it was deprecated

## v1.0.0 (2019)

* Initial release (finished in 2019, officially released 2020-09-16)
