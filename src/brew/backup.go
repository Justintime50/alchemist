package brew

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Justintime50/alchemist/src/general"
)

// Backup backs up your Homebrew instance
func Backup() {
	action := "backup"
	alchemistBackupDir := general.SetupDir(action)
	general.SetupLogging(alchemistBackupDir, action)

	fmt.Println("Alchemist is backing up brew...")

	brewDoctor, brewDoctorErr := general.RunCommand(exec.Command, "brew", []string{"doctor"})
	if brewDoctor != nil {
		log.Printf("brew doctor: %s", brewDoctor)
	} else {
		fmt.Println("Alchemist checked with brew doctor, you need to fix your Homebrew first before it can be backed up!")
		log.Printf("brew update: %s", brewDoctorErr)
		os.Exit(1)
	}

	packageList, _ := general.RunCommand(exec.Command, "brew", []string{"list", "--formula"})
	packages := generatePackageScriptCommands(packageList)

	caskList, _ := general.RunCommand(exec.Command, "brew", []string{"list", "--cask"})
	casks := generateCaskScriptCommands(caskList)

	createScriptFile(packages, alchemistBackupDir+"/restore-brew-packages.sh")
	createScriptFile(casks, alchemistBackupDir+"/restore-brew-casks.sh")

	fmt.Println("Alchemist is finished backing up brew!")
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

// generateCaskScriptCommands creates a script file containing commands to install all your brew casks
func generateCaskScriptCommands(brewCaskList *bytes.Buffer) []string {
	brewCaskListString := brewCaskList.String()
	caskList := strings.Fields(brewCaskListString)

	var caskListArray []string
	caskListArray = append(caskListArray, "#!/bin/sh")
	for i := range caskList {
		caskName := caskList[i]
		brewInstallCommand := fmt.Sprintf("%s %s", "brew install --cask", caskName)
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
