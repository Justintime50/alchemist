package choco

import (
	"bytes"
	"strconv"
	"testing"
)

func TestGeneratePackageScriptCommands(t *testing.T) {
	mockPackageList := bytes.NewBufferString("package1\npackage2")
	wanted := []string{"@echo off", "choco install package1", "choco install package2"}
	output := generatePackageScriptCommands(mockPackageList)
	for i := range output {
		if output[i] != wanted[i] {
			t.Errorf("generatePackageScriptCommands did not properly generate a script with commands. Line %s of the script file got: %s, wanted: %s", strconv.Itoa(i), output[i], wanted[i])
		}
	}
}
