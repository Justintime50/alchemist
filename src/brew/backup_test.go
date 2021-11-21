package brew

import (
	"bytes"
	"strconv"
	"testing"
)

func TestGenerateTapList(t *testing.T) {
	mockPackageList := bytes.NewBufferString("username1/tap1\nusername2/tap2")
	wanted := []string{"# Taps", "tap \"username1/tap1\"", "tap \"username2/tap2\""}
	output := generateTapList(mockPackageList)
	for i := range output {
		if output[i] != wanted[i] {
			t.Errorf("generateTapList did not properly generate a script with commands. Line %s of the script file got: %s, wanted: %s", strconv.Itoa(i), output[i], wanted[i])
		}
	}
}

func TestGeneratePackageList(t *testing.T) {
	mockPackageList := bytes.NewBufferString("package1\npackage2")
	wanted := []string{"# Packages", "brew \"package1\"", "brew \"package2\""}
	output := generatePackageList(mockPackageList)
	for i := range output {
		if output[i] != wanted[i] {
			t.Errorf("generatePackageList did not properly generate a script with commands. Line %s of the script file got: %s, wanted: %s", strconv.Itoa(i), output[i], wanted[i])
		}
	}
}

func TestGenerateCaskList(t *testing.T) {
	mockCaskList := bytes.NewBufferString("cask1\ncask2")
	wanted := []string{"# Casks\ncask_args appdir: \"~/Applications\"", "cask \"cask1\"", "cask \"cask2\""}
	output := generateCaskList(mockCaskList)
	for i := range output {
		if output[i] != wanted[i] {
			t.Errorf("generateCaskList did not properly generate a script with commands. Line %s of the script file got: %s, wanted: %s", strconv.Itoa(i), output[i], wanted[i])
		}
	}
}
