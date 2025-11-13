package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/Justintime50/alchemist/v4/src/brew"
	"github.com/Justintime50/alchemist/v4/src/choco"
)

// main function that accepts CLI args to invoke different functionality
func main() {
	if runtime.GOOS == "windows" {
		chocoUpdate := flag.Bool("update", false, "Update your Chocolatey instance.")
		chocoBackup := flag.Bool("backup", false, "Backup your Chocolatey instance.")
		version := flag.Bool("version", false, "Print the version number of Alchemist.")
		flag.Parse()

		if *version {
			fmt.Println("Alchemist version v4.0.0")
			return
		}
		if *chocoUpdate {
			choco.Update()
			return
		} else if *chocoBackup {
			choco.Backup()
			return
		}
	} else {
		brewUpdate := flag.Bool("update", false, "Update your Homebrew instance.")
		brewBackup := flag.Bool("backup", false, "Backup your Homebrew instance.")
		greedy := flag.Bool("greedy", false, "Greedily updates casks even if they have auto-update capabilities in their respective UIs.")
		version := flag.Bool("version", false, "Print the version number of Alchemist.")
		flag.Parse()

		if *version {
			fmt.Println("Alchemist version v4.0.0")
			return
		}
		if *brewUpdate {
			brew.Update(*greedy)
			return
		} else if *brewBackup {
			brew.Backup()
			return
		}
	}

	fmt.Println("No action taken as no flag was passed.")
	os.Exit(1)
}
