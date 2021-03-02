package choco

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Justintime50/alchemist/src/general"
)

// Backup backs up your Chocolatey instance
func Backup() {
	action := "backup"
	alchemistBackupDir := general.SetupDir(action)
	general.SetupLogging(alchemistBackupDir, action)

	fmt.Println("Alchemist is backing up choco...")

	packageList, _ := general.RunCommand(exec.Command, "choco", []string{"list", "--lo", "--limitoutput", "--id-only"})
	packages := generatePackageScriptCommands(packageList)

	createScriptFile(packages, alchemistBackupDir+"/restore-choco-packages.bat")

	fmt.Println("Alchemist is finished backing up choco!")
}

// generatePackageScriptCommands creates a script file containing commands to install all your choco packages
func generatePackageScriptCommands(chocoPackageList *bytes.Buffer) []string {
	chocoPackageListString := chocoPackageList.String()
	packageList := strings.Fields(chocoPackageListString)

	var packageListArray []string
	packageListArray = append(packageListArray, "@echo off")
	for i := range packageList {
		packageName := packageList[i]
		chocoInstallCommand := fmt.Sprintf("%s %s", "choco install", packageName)
		packageListArray = append(packageListArray, chocoInstallCommand)
	}

	return packageListArray
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
}
