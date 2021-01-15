package update

import (
	"bytes"
	"fmt"
	"github.com/natefinch/lumberjack"
	"log"
	"os"
	"os/exec"
)

// Brew updates your Homebrew instance
func Brew() {
	fmt.Println("Alchemist is updating brew...")
	fmt.Println("This could take some time and may prompt for your password.")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	alchemistUpdateDir := homeDir + "/alchemist/update"
	os.MkdirAll(alchemistUpdateDir, os.ModePerm)

	// TODO: Can we log to a file and print to console in one command for each of these?
	// TODO: Add logging for errors too, only successes get logged currently
	setupLogging(alchemistUpdateDir)

	update := brewUpdate()
	log.Println("brew update: " + update)

	upgrade := brewUpgrade()
	log.Println("brew upgrade: " + upgrade)

	upgradeCasks := brewUpgradeCasks()
	log.Println("brew upgrade --cask: " + upgradeCasks)

	cleanup := brewCleanup()
	log.Println("brew cleanup: " + cleanup)

	doctor := brewDoctor()
	log.Println("brew doctor:" + doctor)
	fmt.Println(doctor)

	fmt.Println("Alchemist is finished updating brew!")
}

// setupLogging sets up logging
func setupLogging(dir string) {
	log.SetOutput(&lumberjack.Logger{
		Filename:   dir + "/alchemist-update.log",
		MaxSize:    1, // megabytes
		MaxBackups: 5,
		MaxAge:     90, // days
	})
}

// brewUpdate updates homebrew
func brewUpdate() string {
	cmd := exec.Command("brew", "update")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error while updating brew: %s.", err))
		os.Exit(1)
	}

	outStrings := out.String()
	fmt.Println("Alchemist has updated brew!")

	return outStrings
}

// brewUpgrade upgrades brew packages
func brewUpgrade() string {
	cmd := exec.Command("brew", "upgrade")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error while upgrading brew packages: %s.", err))
		os.Exit(1)
	}

	outStrings := out.String()
	fmt.Println("Alchemist has upgraded brew packages!")

	return outStrings
}

// brewUpgradeCasks upgrades brew packages
func brewUpgradeCasks() string {
	cmd := exec.Command("brew", "upgrade", "--cask")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error while upgrading brew casks: %s.", err))
		os.Exit(1)
	}

	outStrings := out.String()
	fmt.Println("Alchemist has upgraded brew casks!")

	return outStrings
}

// brewCleanup cleans homebrew
func brewCleanup() string {
	cmd := exec.Command("brew", "cleanup")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error while cleaning brew: %s.", err))
		os.Exit(1)
	}

	outStrings := out.String()
	fmt.Println("Alchemist has cleaned up brew!")

	return outStrings
}

// brewDoctor checks that the homebrew instance is healthy
func brewDoctor() string {
	cmd := exec.Command("brew", "doctor")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error while checking with brew doctor: %s.", err))
		os.Exit(1)
	}

	outStrings := out.String()
	fmt.Println("Alchemist has checked with brew doctor:")

	return outStrings
}
