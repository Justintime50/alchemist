package update

import (
	"fmt"
	"os"
	"os/exec"
)

// UpdateBrew updates your Homebrew instance
func UpdateBrew() {
	fmt.Printf("%s\n", "Alchemist is updating brew...")
	fmt.Printf("%s\n", "This could take some time and may prompt for your password.")

	updateErr := exec.Command("brew", "update").Run()
	if updateErr != nil {
		fmt.Printf("%s\n", fmt.Sprintf("Error while updating brew: %s.", updateErr))
		os.Exit(1)
	}
	fmt.Printf("%s\n", "Alchemist has updated brew!")

	upgradeErr := exec.Command("brew", "upgrade").Run()
	if upgradeErr != nil {
		fmt.Printf("%s\n", fmt.Sprintf("Error while upgrading brew packages: %s.", upgradeErr))
		os.Exit(1)
	}
	fmt.Printf("%s\n", "Alchemist has upgraded brew packages!")

	upgradeCaskErr := exec.Command("brew", "upgrade", "--cask").Run()
	if upgradeCaskErr != nil {
		fmt.Printf("%s\n", fmt.Sprintf("Error while upgrading brew casks: %s.", upgradeCaskErr))
		os.Exit(1)
	}
	fmt.Printf("%s\n", "Alchemist has upgraded brew casks!")

	cleanupErr := exec.Command("brew", "cleanup").Run()
	if cleanupErr != nil {
		fmt.Printf("%s\n", fmt.Sprintf("Error while cleaning brew: %s.", cleanupErr))
		os.Exit(1)
	}
	fmt.Printf("%s\n", "Alchemist has cleaned up brew!")

	doctorErr := exec.Command("brew", "doctor").Run()
	if doctorErr != nil {
		fmt.Printf("%s\n", fmt.Sprintf("Error while checking with brew doctor: %s.", doctorErr))
		os.Exit(1)
	}
	fmt.Printf("%s\n", "Alchemist has checked with brew doctor!")
	// TODO: Display the output of `brew doctor`

	fmt.Printf("%s\n", "Alchemist is finished updating brew!")
}
