package general

import (
	"os"
	"testing"

	"github.com/justintime50/mockcmd"
)

func TestSetupDir(t *testing.T) {
	setupDir := SetupDir("mock-directory")
	if setupDir == "/alchemist/mock-directory" {
		t.Errorf("setupDir did not concatenate the homedir to the app path, got: %s", setupDir)
	}
	// Tear down the mocked directory
	// TODO: in a perfect world, we'd actually mock this and not commit to disk and remove
	_ = os.Remove(setupDir)
}

func TestRunCommandSuccess(t *testing.T) {
	stdout, err := RunCommand(mockcmd.MockExecSuccess, "mock-command", []string{"mock", "args"})
	mockcmd.Success(t, stdout, err)
}

func TestRunCommandFailure(t *testing.T) {
	_, err := RunCommand(mockcmd.MockExecFailure, "mock-command", []string{"mock", "args"})
	mockcmd.Fail(t, err)
}

// The following two functions are required by `mockcmd` for all mocked cmd tests
func TestMockProcessSuccess(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	_, _ = os.Stdout.WriteString("mocked Stdout")
	os.Exit(0)
}

func TestMockProcessFailure(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	os.Exit(1)
}
