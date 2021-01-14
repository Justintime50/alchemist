package backup

import (
	"fmt"
	// "os"
	"bytes"
	"os/exec"
	"strings"
)

// BackupBrew backs up your Homebrew instance
func BackupBrew() {
	fmt.Printf("%s\n", fmt.Sprintf("Alchemist is backing up brew..."))

	BackupPackages()
	BackupCasks()

	// fmt.Printf("%v", test)
	// outfile, err := os.Create("./out.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer outfile.Close()
	// cmd.Wait()

	fmt.Printf("%s\n", "Alchemist is finished backing up brew!")
}

// BackupPackages creates a script file containing commands to install all your brew packages
func BackupPackages() {
	cmd := exec.Command("brew", "list", "--formula")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	outStrings := out.String()
	packageList := strings.Fields(outStrings)

	fmt.Printf("%s\n", "#!/bin/sh")
	for i := range packageList {
		packageName := packageList[i]
		fmt.Printf("%s\n", fmt.Sprintf("%s %s", "brew install", packageName))
	}

	// TODO: Create a new array in memory and save it to the script file
}

// BackupCasks creates a script file containing commands to install all your brew casks
func BackupCasks() {
	cmd := exec.Command("brew", "list", "--cask")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	outStrings := out.String()
	caskList := strings.Fields(outStrings)

	fmt.Printf("%s\n", "#!/bin/sh")
	for i := range caskList {
		caskName := caskList[i]
		fmt.Printf("%s\n", fmt.Sprintf("%s %s", "brew cask install", caskName))
	}

	// TODO: Create a new array in memory and save it to the script file
}
