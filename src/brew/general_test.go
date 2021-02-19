package brew

import (
	"os"
	"testing"

	"github.com/justintime50/mockcmd/mockcmd"
)

func TestSetupDir(t *testing.T) {
	setupDir := setupDir("mock-directory")
	if setupDir == "/alchemist/backup" {
		t.Errorf("setupDir did not concatenate the homedir to the app path, got: %s", setupDir)
	}
}

func TestRunCommandSuccess(t *testing.T) {
	stdout, err := runCommand(mockcmd.MockExecSuccess, "mock-command", []string{"mock", "args"})
	mockcmd.Success(t, stdout, err)
}

func TestRunCommandFailure(t *testing.T) {
	_, err := runCommand(mockcmd.MockExecFailure, "mock-command", []string{"mock", "args"})
	mockcmd.Fail(t, err)
}

// The following two functions are required by `mockcmd` for all mocked cmd tests
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
