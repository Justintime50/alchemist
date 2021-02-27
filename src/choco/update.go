package choco

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Update updates your Chocolatey instance
func Update() {
	action := "update"
	alchemistUpdateDir := setupDir(action)
	setupLogging(alchemistUpdateDir, action)

	fmt.Println("Alchemist is updating choco...")
	fmt.Println("This could take some time.")

	update, updateErr := runCommand(exec.Command, "choco", []string{"upgrade", "all", "-y"})
	if update != nil {
		fmt.Println("Alchemist updated choco!")
		log.Printf("choco update: %s", update)
	} else {
		fmt.Println("Alchemist could not update choco, please see logs for details.")
		log.Printf("choco update: %s", updateErr)
		os.Exit(1)
	}

	fmt.Println("Alchemist is finished updating choco!")
}
