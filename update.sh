#!/bin/bash

DATE=$(date +'%m-%d-%Y')
LOG_LIFE="30" 
LOCATION="$HOME/brew-update"

mkdir -p "$LOCATION"

{
    brew update
    brew upgrade
    brew cask upgrade
    brew cleanup
    brew doctor
} 2>&1 | tee "$LOCATION"/logs/"$DATE".log

find "$LOCATION" -mindepth 1 -mtime +"$LOG_LIFE" -delete
