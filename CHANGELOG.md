# CHANGELOG

## v2.2.0 (2021-03-01)

* Add support for `choco` on `Windows`
* Alchemist now detects which OS you're running on and changes its logic accordingly

## v2.1.0 (2021-02-18)

* Completely overhauled the app's modules by combining update/backup modules into a new `brew` module in preperation for adding `Choco` support on Windows
* Cut out a lot of duplicate code
* Bug fix for the `brew cask` install command on backup scripts
* Bug fix for printing `nil` to console instead of stderr (closes #3)

## v2.0.0 (2021-01-18)

* Renamed `brew-update` to `alchemist`
* Consolidated `brew-backup` tool into Alchemist, now you can specify the functionality by passing a flag (eg: `--update`, `--backup`)
* Rewrote the tool from shell scripts into Golang
* Removed the reliance on system tools like `sed` for much safer and compatible functionality
* Added proper logging framework with automatic rollover (90 day log life, 1mb file size, 5 logs max)
* Added unit tests
* Auto releasing via `goreleaser`

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
