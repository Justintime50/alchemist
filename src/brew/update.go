package brew

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Justintime50/alchemist/src/general"
)

// Update updates your Homebrew instance
func Update() {
	action := "update"
	alchemistUpdateDir := general.SetupDir(action)
	general.SetupLogging(alchemistUpdateDir, action)

	fmt.Println("Alchemist is updating brew...")
	fmt.Println("This could take some time and may prompt for your password.")

	update, updateErr := general.RunCommand(exec.Command, "brew", []string{"update"})
	if update != nil {
		fmt.Println("Alchemist updated brew!")
		log.Printf("brew update: %s", update)
	} else {
		fmt.Println("Alchemist could not update brew, please see logs for details.")
		log.Printf("brew update: %s", updateErr)
		os.Exit(1)
	}

	fmt.Println("Alchemist is upgrading brew packages...")
	upgrade, upgradeErr := general.RunCommand(exec.Command, "brew", []string{"upgrade"})
	if upgrade != nil {
		fmt.Println("Alchemist upgraded brew packages!")
		log.Printf("brew upgrade: %s", upgrade)
	} else {
		fmt.Println("Alchemist could not upgrade brew packages, please see logs for details.")
		log.Printf("brew upgrade: %s", upgradeErr)
		os.Exit(1)
	}

	fmt.Println("Alchemist is upgrading brew casks...")
	upgradeCasks, upgradeCasksErr := general.RunCommand(exec.Command, "brew", []string{"upgrade", "--cask"})
	if upgradeCasks != nil {
		fmt.Println("Alchemist upgraded brew casks!")
		log.Printf("brew upgrade --cask: %s", upgradeCasks)
	} else {
		fmt.Println("Alchemist could not upgrade brew casks, please see logs for details.")
		log.Printf("brew upgrade --cask: %s", upgradeCasksErr)
		os.Exit(1)
	}

	fmt.Println("Alchemist is cleaning up brew...")
	cleanup, cleanupErr := general.RunCommand(exec.Command, "brew", []string{"cleanup"})
	if cleanup != nil {
		fmt.Println("Alchemist cleaned brew!")
		log.Printf("brew cleanup: %s", cleanup)
	} else {
		fmt.Println("Alchemist could not clean brew, please see logs for details.")
		log.Printf("brew cleanup: %s", cleanupErr)
		os.Exit(1)
	}

	fmt.Println("Alchemist is checking with brew doctor...")
	doctor, doctorErr := general.RunCommand(exec.Command, "brew", []string{"doctor"})
	if cleanup != nil {
		fmt.Println("Alchemist doctored brew!")
		log.Printf("brew doctor: %s", doctor)
	} else {
		fmt.Println("Alchemist could not doctor brew, please see logs for details.")
		log.Printf("brew doctor: %s", doctorErr)
		os.Exit(1)
	}

	fmt.Println("Alchemist is finished updating brew!")
}
