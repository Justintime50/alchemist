package update

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/justintime50/mockcmd/mockcmd"
	"github.com/natefinch/lumberjack"
)

// Brew updates your Homebrew instance
func Brew() {
	fmt.Println("Alchemist is updating brew...")
	fmt.Println("This could take some time and may prompt for your password.")

	alchemistUpdateDir := setupDir()

	// TODO: Can we log to a file and print to console in one command for each of these?
	setupLogging(alchemistUpdateDir)

	update, _ := brewUpdate(exec.Command)
	fmt.Println("Alchemist has updated brew!")
	log.Printf("brew update: %s", update.String())

	upgrade, _ := brewUpgrade(exec.Command)
	fmt.Println("Alchemist has upgraded brew packages!")
	log.Printf("brew upgrade: %s", upgrade.String())

	upgradeCasks, _ := brewUpgradeCasks(exec.Command)
	fmt.Println("Alchemist has upgraded brew casks!")
	log.Printf("brew upgrade --cask: %s", upgradeCasks.String())

	cleanup, _ := brewCleanup(exec.Command)
	fmt.Println("Alchemist has cleaned up brew!")
	log.Printf("brew cleanup: %s", cleanup.String())

	doctor, _ := brewDoctor(exec.Command)
	fmt.Println("Alchemist has checked with brew doctor:")
	log.Printf("brew doctor: %s", doctor.String())
	fmt.Println(doctor)

	fmt.Println("Alchemist is finished updating brew!")
}

// setupDir gets the user's home dir and appends the alchemist dir to it
func setupDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	alchemistUpdateDir := homeDir + "/alchemist/update"
	os.MkdirAll(alchemistUpdateDir, os.ModePerm)
	return alchemistUpdateDir
}

// setupLogging sets up logging
func setupLogging(dir string) string {
	logPath := dir + "/alchemist-update.log"
	log.SetOutput(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    1, // megabytes
		MaxBackups: 5,
		MaxAge:     90, // days
	})
	return logPath
}

// brewUpdate updates homebrew
func brewUpdate(cmdContext mockcmd.ExecContext) (*bytes.Buffer, error) {
	cmd := cmdContext("brew", "update")
	var outb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &outb
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error while updating brew: %s.", err))
		return nil, err
	}

	return &outb, nil
}

// brewUpgrade upgrades brew packages
func brewUpgrade(cmdContext mockcmd.ExecContext) (*bytes.Buffer, error) {
	cmd := cmdContext("brew", "upgrade")
	var outb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &outb
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error while upgrading brew packages: %s.", err))
		return nil, err
	}

	return &outb, nil
}

// brewUpgradeCasks upgrades brew packages
func brewUpgradeCasks(cmdContext mockcmd.ExecContext) (*bytes.Buffer, error) {
	cmd := cmdContext("brew", "upgrade", "--cask")
	var outb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &outb
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error while upgrading brew casks: %s.", err))
		return nil, err
	}

	return &outb, nil
}

// brewCleanup cleans homebrew
func brewCleanup(cmdContext mockcmd.ExecContext) (*bytes.Buffer, error) {
	cmd := cmdContext("brew", "cleanup")
	var outb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &outb
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error while cleaning brew: %s.", err))
		return nil, err
	}

	return &outb, nil
}

// brewDoctor checks that the homebrew instance is healthy
func brewDoctor(cmdContext mockcmd.ExecContext) (*bytes.Buffer, error) {
	cmd := cmdContext("brew", "doctor")
	var outb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &outb
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error while checking with brew doctor: %s.", err))
		return nil, err
	}

	return &outb, nil
}
