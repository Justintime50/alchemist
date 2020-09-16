#!/bin/bash

DATE=$(date +'%m-%d-%Y')
LOG_LIFE="90" 
LOCATION="$HOME/brew-update"

mkdir -p "$LOCATION"
mkdir -p "$LOCATION"/logs

{
    brew update
    brew upgrade
    brew upgrade --cask
    brew cleanup
    brew doctor
} 2>&1 | tee "$LOCATION"/logs/"$DATE".log

find "$LOCATION"/logs -mindepth 1 -mtime +"$LOG_LIFE" -delete
