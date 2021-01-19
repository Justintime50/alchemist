package backup

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/justintime50/mockcmd/mockcmd"
)

// Brew backs up your Homebrew instance
func Brew() {
	fmt.Println("Alchemist is backing up brew...")

	packageList, _ := retrieveBrewList(exec.Command)
	packages := generatePackageScriptCommands(packageList)

	caskList, _ := retrieveBrewCaskList(exec.Command)
	casks := generateCaskScriptCommands(caskList)

	alchemistBackupDir := setupDir()
	createScriptFile(packages, alchemistBackupDir+"/restore-brew-packages.sh")
	createScriptFile(casks, alchemistBackupDir+"/restore-brew-casks.sh")

	fmt.Println("Alchemist is finished backing up brew!")
}

// setupDir gets the user's home dir and appends the alchemist dir to it
func setupDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	alchemistBackupDir := homeDir + "/alchemist/backup"
	_ = os.MkdirAll(alchemistBackupDir, os.ModePerm)
	return alchemistBackupDir
}

// retrieveBrewList retrieves the list of brew packages from Homebrew
func retrieveBrewList(cmdContext mockcmd.ExecContext) (*bytes.Buffer, error) {
	cmd := cmdContext("brew", "list", "--formula")
	var outb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &outb
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error while retrieving brew list: %s.\n", err)
		return nil, err
	}

	return &outb, nil
}

// generatePackageScriptCommands creates a script file containing commands to install all your brew packages
func generatePackageScriptCommands(brewPackageList *bytes.Buffer) []string {
	brewPackageListString := brewPackageList.String()
	packageList := strings.Fields(brewPackageListString)

	var packageListArray []string
	packageListArray = append(packageListArray, "#!/bin/sh")
	for i := range packageList {
		packageName := packageList[i]
		brewInstallCommand := fmt.Sprintf("%s %s", "brew install", packageName)
		packageListArray = append(packageListArray, brewInstallCommand)
	}

	return packageListArray
}

// retrieveBrewCaskList retrieves the list of brew casks from Homebrew
func retrieveBrewCaskList(cmdContext mockcmd.ExecContext) (*bytes.Buffer, error) {
	cmd := cmdContext("brew", "list", "--cask")
	var outb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &outb
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error while retrieving brew cask list: %s.\n", err)
		return nil, err
	}

	return &outb, nil
}

// generateCaskScriptCommands creates a script file containing commands to install all your brew casks
func generateCaskScriptCommands(brewCaskList *bytes.Buffer) []string {
	brewCaskListString := brewCaskList.String()
	caskList := strings.Fields(brewCaskListString)

	var caskListArray []string
	caskListArray = append(caskListArray, "#!/bin/sh")
	for i := range caskList {
		caskName := caskList[i]
		brewInstallCommand := fmt.Sprintf("%s %s", "brew cask install", caskName)
		caskListArray = append(caskListArray, brewInstallCommand)
	}

	return caskListArray
}

// createScriptFile creates a script file from an array of commands
func createScriptFile(commands []string, filename string) {
	packageFile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer packageFile.Close()

	for i := range commands {
		command := commands[i]
		_, _ = packageFile.WriteString(fmt.Sprintf("%s\n", command))
	}

	err = exec.Command("chmod", "+x", filename).Run()
	if err != nil {
		os.Exit(1)
	}
}
