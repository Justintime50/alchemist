package main

import (
	"flag"
	"fmt"
	"github.com/Justintime50/alchemist/src/backup"
	"github.com/Justintime50/alchemist/src/update"
)

func main() {
	// main function that accepts CLI args to invoke different functionality
	brewUpdate := flag.Bool("update", false, "Update your Homebrew instance.")
	brewBackup := flag.Bool("backup", false, "Backup your Homebrew instance.")
	flag.Parse()

	if *brewUpdate {
		update.UpdateBrew()
		return
	} else if *brewBackup {
		backup.BackupBrew()
		return
	}

	// TODO: Add logging

	fmt.Printf("%s\n", "No action taken as no flag was passed.")
}
