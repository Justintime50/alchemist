package backup

import (
	"bytes"
	"os"
	"strconv"
	"testing"

	"github.com/justintime50/mockcmd/mockcmd"
)

// TODO: Mock and test the `createScriptFile()` function

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
	wanted := []string{"#!/bin/sh", "brew cask install cask1", "brew cask install cask2"}
	output := generateCaskScriptCommands(mockCaskList)
	for i := range output {
		if output[i] != wanted[i] {
			t.Errorf("generateCaskScriptCommands did not properly generate a script with commands. Line %s of the script file got: %s, wanted: %s", strconv.Itoa(i), output[i], wanted[i])
		}
	}
}

func TestRetrieveBrewListSuccess(t *testing.T) {
	stdout, err := retrieveBrewList(mockcmd.MockExecSuccess)
	mockcmd.Success(t, stdout, err)
}

func TestRetrieveBrewListFailure(t *testing.T) {
	_, err := retrieveBrewList(mockcmd.MockExecFailure)
	mockcmd.Fail(t, err)
}

func TestRetrieveBrewCaskListSuccess(t *testing.T) {
	stdout, err := retrieveBrewCaskList(mockcmd.MockExecSuccess)
	mockcmd.Success(t, stdout, err)
}

func TestRetrieveBrewCaskListFailure(t *testing.T) {
	_, err := retrieveBrewCaskList(mockcmd.MockExecFailure)
	mockcmd.Fail(t, err)
}

// The following two functions are required by `mockcmd`
func TestMockProcessSuccess(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	os.Stdout.WriteString("mocked Stdout")
	os.Exit(0)
}

func TestMockProcessFailure(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	os.Exit(1)
}
