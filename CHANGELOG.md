# CHANGELOG

## v3.3.1 (2024-11-22)

- Only upgrade Casks on macOS since they aren't supported on Linux

## v3.3.0 (2024-08-22)

- Adds support for Go `1.22` and `1.23`

## v3.2.0 (2022-05-21)

- Adds a `--greedy` CLI flag so casks can be updated even if they have a UI "auto-update" feature (helps keep Homebrew app versions synced better with what they actually are)

## v3.1.3 (2022-03-23)

- Corrects project namespace by appending `/v3`
- Bumps dependencies

## v3.1.2 (2021-12-09)

- Adds back support for Go 1.16 which was previously removed in `3.1.0`

## v3.1.1 (2021-11-21)

- No longer require `sha checksum` on cask installation as some casks didn't include them which would lead to items on your backup list not getting installed

## v3.1.0 (2021-09-21)

- Drops support for Go `1.15`
- Targets Go `1.17`
- Update dependencies

## v3.0.0 (2021-05-04)

- Replaces `restore-brew-packages.sh` and `restore-brew-casks.sh` with a single `Brewfile` which also now includes all installed `taps`
- Adds `--force` flag to force Brew backups even if errors are present in `brew doctor`

## v2.2.1 (2021-04-09)

- Fixes a bug where we left a mocked directory on disk (closes #6)
- Fixes a bug in the `TestSetupDir` test that would assert the wrong value

## v2.2.0 (2021-03-01)

- Add support for `choco` on `Windows`
- Alchemist now detects which OS you're running on and changes its logic accordingly

## v2.1.0 (2021-02-18)

- Completely overhauled the app's modules by combining update/backup modules into a new `brew` module in preperation for adding `Choco` support on Windows
- Cut out a lot of duplicate code
- Bug fix for the `brew cask` install command on backup scripts
- Bug fix for printing `nil` to console instead of stderr (closes #3)

## v2.0.0 (2021-01-18)

- Renamed `brew-update` to `alchemist`
- Consolidated `brew-backup` tool into Alchemist, now you can specify the functionality by passing a flag (eg: `--update`, `--backup`)
- Rewrote the tool from shell scripts into Golang
- Removed the reliance on system tools like `sed` for much safer and compatible functionality
- Added proper logging framework with automatic rollover (90 day log life, 1mb file size, 5 logs max)
- Added unit tests
- Auto releasing via `goreleaser`

## v1.2.1 (2021-01-09)

- Switching from Travis-CI to GitHub Actions
- Using `homebrew-releaser` for automated releasing

## v1.2.0 (2020-10-05)

- Refactored code to use functions
- Added parameters to allow for a custom log path and log life in days

## v1.1.0 (2020-09-16)

- Now installable via Homebrew
- Removed launch agent
- Replaced `brew cask upgrade` with `brew upgrade --cask` as it was deprecated

## v1.0.0 (2019)

- Initial release (finished in 2019, officially released 2020-09-16)
