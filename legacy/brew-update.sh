#!/bin/bash

main() {
    LOCATION=${1:-"$HOME/brew-update"}
    LOG_LIFE=${2:-"90"}
    
    make_directories
    brew_update
    clean_logs
}

make_directories() {
    # Create project directories if they do not exist
    mkdir -p "$LOCATION"
}

brew_update() {
    # Update, upgrade, clean, and doctor your Homebrew instance
    local date
    date=$(date +'%Y-%m-%d')
    {
        brew update
        brew upgrade
        brew upgrade --cask
        brew cleanup
        brew doctor
    } 2>&1 | tee "$LOCATION"/"$date".log
}

clean_logs() {
    # If logs are older than the log life, delete them
    find "$LOCATION"/logs -mindepth 1 -mtime +"$LOG_LIFE" -delete
}

main "$1" "$2"
