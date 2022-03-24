package brew

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Justintime50/alchemist/v3/src/general"
)

// Backup backs up your Homebrew instance
func Backup(force bool) {
	action := "backup"
	alchemistBackupDir := general.SetupDir(action)
	general.SetupLogging(alchemistBackupDir, action)

	fmt.Println("Alchemist is backing up brew...")

	brewDoctor, brewDoctorErr := general.RunCommand(exec.Command, "brew", []string{"doctor"})
	if brewDoctor != nil || force {
		log.Printf("brew doctor: %s", brewDoctor)
	} else {
		fmt.Println("Alchemist checked with brew doctor, you need to fix your Homebrew instance before it can be backed up!")
		log.Printf("brew doctor: %s", brewDoctorErr)
		os.Exit(1)
	}

	tapList, _ := general.RunCommand(exec.Command, "brew", []string{"tap"})
	taps := generateTapList(tapList)

	packageList, _ := general.RunCommand(exec.Command, "brew", []string{"list", "--formula"})
	packages := generatePackageList(packageList)

	caskList, _ := general.RunCommand(exec.Command, "brew", []string{"list", "--cask"})
	casks := generateCaskList(caskList)

	brewfileContent := append(taps, packages...)
	brewfileContent = append(brewfileContent, casks...)

	createBrewfile(brewfileContent, alchemistBackupDir+"/Brewfile")

	fmt.Println("Alchemist is finished backing up brew!")
}

// generateTapList creates a list of brew taps
func generateTapList(brewTapList *bytes.Buffer) []string {
	brewTapListString := brewTapList.String()
	tapList := strings.Fields(brewTapListString)

	var tapListArray []string
	tapListArray = append(tapListArray, "# Taps")
	for i := range tapList {
		tapName := tapList[i]
		brewInstallCommand := fmt.Sprintf("%s \"%s\"", "tap", tapName)
		tapListArray = append(tapListArray, brewInstallCommand)
	}

	return tapListArray
}

// generatePackageList creates a list of brew packages
func generatePackageList(brewPackageList *bytes.Buffer) []string {
	brewPackageListString := brewPackageList.String()
	packageList := strings.Fields(brewPackageListString)

	var packageListArray []string
	packageListArray = append(packageListArray, "# Packages")
	for i := range packageList {
		packageName := packageList[i]
		brewInstallCommand := fmt.Sprintf("%s \"%s\"", "brew", packageName)
		packageListArray = append(packageListArray, brewInstallCommand)
	}

	return packageListArray
}

// generateCaskList creates a list of brew casks
func generateCaskList(brewCaskList *bytes.Buffer) []string {
	brewCaskListString := brewCaskList.String()
	caskList := strings.Fields(brewCaskListString)

	var caskListArray []string
	caskListArray = append(caskListArray, "# Casks\ncask_args appdir: \"~/Applications\"")
	for i := range caskList {
		caskName := caskList[i]
		brewInstallCommand := fmt.Sprintf("%s \"%s\"", "cask", caskName)
		caskListArray = append(caskListArray, brewInstallCommand)
	}

	return caskListArray
}

// createBrewfile creates a Brewfile from an array of commands
func createBrewfile(commands []string, filename string) {
	brewfile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer brewfile.Close()

	for i := range commands {
		command := commands[i]
		_, _ = brewfile.WriteString(fmt.Sprintf("%s\n", command))
	}

	err = exec.Command("chmod", "+x", filename).Run()
	if err != nil {
		os.Exit(1)
	}
}
