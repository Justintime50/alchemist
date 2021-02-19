package brew

import (
	"bytes"
	"strconv"
	"testing"
)

func TestGeneratePackageScriptCommands(t *testing.T) {
	mockPackageList := bytes.NewBufferString("package1\npackage2")
	wanted := []string{"#!/bin/sh", "brew install package1", "brew install package2"}
	output := generatePackageScriptCommands(mockPackageList)
	for i := range output {
		if output[i] != wanted[i] {
			t.Errorf("generatePackageScriptCommands did not properly generate a script with commands. Line %s of the script file got: %s, wanted: %s", strconv.Itoa(i), output[i], wanted[i])
		}
	}
}

func TestGenerateCaskScriptCommands(t *testing.T) {
	mockCaskList := bytes.NewBufferString("cask1\ncask2")
	wanted := []string{"#!/bin/sh", "brew install --cask cask1", "brew install --cask cask2"}
	output := generateCaskScriptCommands(mockCaskList)
	for i := range output {
		if output[i] != wanted[i] {
			t.Errorf("generateCaskScriptCommands did not properly generate a script with commands. Line %s of the script file got: %s, wanted: %s", strconv.Itoa(i), output[i], wanted[i])
		}
	}
}
