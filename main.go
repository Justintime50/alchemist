package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/Justintime50/alchemist/v3/src/brew"
	"github.com/Justintime50/alchemist/v3/src/choco"
)

// main function that accepts CLI args to invoke different functionality
func main() {
	if runtime.GOOS == "windows" {
		chocoUpdate := flag.Bool("update", false, "Update your Chocolatey instance.")
		chocoBackup := flag.Bool("backup", false, "Backup your Chocolatey instance.")
		flag.Parse()

		if *chocoUpdate {
			choco.Update()
			return
		} else if *chocoBackup {
			choco.Backup()
			return
		}

		fmt.Println("No action taken as no flag was passed.")
		os.Exit(1)
	} else {
		brewUpdate := flag.Bool("update", false, "Update your Homebrew instance.")
		brewBackup := flag.Bool("backup", false, "Backup your Homebrew instance.")
		force := flag.Bool("force", false, "Forces actions such as backing up even when there are errors.")
		greedy := flag.Bool("greedy", false, "Force updates to casks that have auto-update capabilities in their respective UIs.")
		flag.Parse()

		if *brewUpdate {
			brew.Update(*greedy)
			return
		} else if *brewBackup {
			brew.Backup(*force)
			return
		}

		fmt.Println("No action taken as no flag was passed.")
		os.Exit(1)
	}
}
