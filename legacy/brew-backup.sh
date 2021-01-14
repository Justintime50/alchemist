#!/bin/bash

# Backup your Homebrew packages and casks to installable scripts.

main() {
    BREW_BACKUP_LOCATION=${1:-"$HOME/brew-backup"}
    PACKAGE_SCRIPT_LOCATION="$BREW_BACKUP_LOCATION/restore-brew-packages.sh"
    CASK_SCRIPT_LOCATION="$BREW_BACKUP_LOCATION/restore-brew-casks.sh"
    
    echo -e "Backing up Homebrew packages and casks...\n"
    setup_location_folder
    create_restore_scripts
    brew_backup
    make_packages_installable

    make_script_executable "$PACKAGE_SCRIPT_LOCATION"
    make_script_executable "$CASK_SCRIPT_LOCATION"

    echo "Homebrew package restore script saved to: $PACKAGE_SCRIPT_LOCATION"
    echo "Homebrew cask restore script saved to: $CASK_SCRIPT_LOCATION"
}

setup_location_folder() {
    # If the location folder does not exists, set it up
    mkdir -p "$BREW_BACKUP_LOCATION"
}

make_script_executable() {
    # Make script executable
    chmod 755 "$1"
}

create_restore_scripts() {
    # Create the empty restore scripts with a shebang
    local timestamp
    timestamp=$(date +%Y-%m-%d:%H-%M-%S)
    echo -e "#!/bin/bash\n# Backed up: $timestamp" > "$PACKAGE_SCRIPT_LOCATION"
    echo -e "#!/bin/bash\n# Backed up: $timestamp" > "$CASK_SCRIPT_LOCATION"
}

brew_backup() {
    # Populate Brew data to a script file
    brew list --formula >> "$PACKAGE_SCRIPT_LOCATION"
    brew list --cask >> "$CASK_SCRIPT_LOCATION"
}

make_packages_installable() {
    # Make scripts installable by appending the brew commands to the package names
    # TODO: This usage of "sed" may not work on all flavors of Unix, replace if possible
    sed -i "" '3,$s/^/brew install /' "$PACKAGE_SCRIPT_LOCATION"
    sed -i "" '3,$s/^/brew cask install /' "$CASK_SCRIPT_LOCATION"
}

main "$1"
