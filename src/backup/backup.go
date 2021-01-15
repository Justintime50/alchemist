package backup

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Brew backs up your Homebrew instance
func Brew() {
	fmt.Println("Alchemist is backing up brew...")

	packages := generatePackageScriptCommands()
	casks := generateCaskScriptCommands()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	alchemistBackupDir := homeDir + "/alchemist/backup"
	os.MkdirAll(alchemistBackupDir, os.ModePerm)
	createScriptFile(packages, alchemistBackupDir+"/restore-brew-packages.sh")
	createScriptFile(casks, alchemistBackupDir+"/restore-brew-casks.sh")

	fmt.Println("Alchemist is finished backing up brew!")
}

// generatePackageScriptCommands creates a script file containing commands to install all your brew packages
func generatePackageScriptCommands() []string {
	cmd := exec.Command("brew", "list", "--formula")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}

	outStrings := out.String()
	packageList := strings.Fields(outStrings)

	var packageListArray []string
	packageListArray = append(packageListArray, "#!/bin/sh")
	for i := range packageList {
		packageName := packageList[i]
		brewInstallCommand := fmt.Sprintf("%s %s", "brew install", packageName)
		packageListArray = append(packageListArray, brewInstallCommand)
	}

	return packageListArray
}

// generateCaskScriptCommands creates a script file containing commands to install all your brew casks
func generateCaskScriptCommands() []string {
	cmd := exec.Command("brew", "list", "--cask")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}

	outStrings := out.String()
	caskList := strings.Fields(outStrings)

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
		packageFile.WriteString(fmt.Sprintf("%s\n", command))
	}

	err = exec.Command("chmod", "+x", filename).Run()
	if err != nil {
		os.Exit(1)
	}
}
