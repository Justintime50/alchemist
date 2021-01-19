package main

import (
	"flag"
	"fmt"
	"github.com/Justintime50/alchemist/src/backup"
	"github.com/Justintime50/alchemist/src/update"
	"os"
)

// main function that accepts CLI args to invoke different functionality
func main() {
	brewUpdate := flag.Bool("update", false, "Update your Homebrew instance.")
	brewBackup := flag.Bool("backup", false, "Backup your Homebrew instance.")
	flag.Parse()

	if *brewUpdate {
		update.Brew()
		return
	} else if *brewBackup {
		backup.Brew()
		return
	}

	fmt.Println("No action taken as no flag was passed.")
	os.Exit(1)
}
