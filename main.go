package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Justintime50/alchemist/src/brew"
)

// main function that accepts CLI args to invoke different functionality
func main() {
	brewUpdate := flag.Bool("update", false, "Update your Homebrew instance.")
	brewBackup := flag.Bool("backup", false, "Backup your Homebrew instance.")
	flag.Parse()

	// TODO: Check what OS is running and split macOS/Linux and Windows here
	if *brewUpdate {
		brew.Update()
		return
	} else if *brewBackup {
		brew.Backup()
		return
	}

	fmt.Println("No action taken as no flag was passed.")
	os.Exit(1)
}
