package update

import (
	"os"
	"testing"

	"github.com/justintime50/mockcmd/mockcmd"
)

// TODO: Test `Brew()` and that each function is called correctly

func TestSetupDir(t *testing.T) {
	setupDir := setupDir()
	if setupDir == "/alchemist/update" {
		t.Errorf("setupDir did not concatenate the homedir to the app path, got: %s", setupDir)
	}
}

func TestSetupLogging(t *testing.T) {
	setupLogging := setupLogging("mock-dir")
	logPath := "mock-dir/alchemist-update.log"
	if setupLogging != logPath {
		t.Errorf("setupLogging did not build a proper logPath, got: %s, want: %s.", setupLogging, logPath)
	}
}

func TestBrewUpdateSuccess(t *testing.T) {
	stdout, err := brewUpdate(mockcmd.MockExecSuccess)
	mockcmd.Success(t, stdout, err)
}

func TestBrewUpdateFailure(t *testing.T) {
	_, err := brewUpdate(mockcmd.MockExecFailure)
	mockcmd.Fail(t, err)
}

func TestBrewUpgradeSuccess(t *testing.T) {
	stdout, err := brewUpgrade(mockcmd.MockExecSuccess)
	mockcmd.Success(t, stdout, err)
}

func TestBrewUpgradeFailure(t *testing.T) {
	_, err := brewUpgrade(mockcmd.MockExecFailure)
	mockcmd.Fail(t, err)
}

func TestBrewUpgradeCasksSuccess(t *testing.T) {
	stdout, err := brewUpgradeCasks(mockcmd.MockExecSuccess)
	mockcmd.Success(t, stdout, err)
}

func TestBrewUpgradeCasksFailure(t *testing.T) {
	_, err := brewUpgradeCasks(mockcmd.MockExecFailure)
	mockcmd.Fail(t, err)
}

func TestBrewCleanupSuccess(t *testing.T) {
	stdout, err := brewCleanup(mockcmd.MockExecSuccess)
	mockcmd.Success(t, stdout, err)
}

func TestBrewCleanupFailure(t *testing.T) {
	_, err := brewCleanup(mockcmd.MockExecFailure)
	mockcmd.Fail(t, err)
}

func TestBrewDoctorSuccess(t *testing.T) {
	stdout, err := brewDoctor(mockcmd.MockExecSuccess)
	mockcmd.Success(t, stdout, err)
}

func TestBrewDoctorFailure(t *testing.T) {
	_, err := brewDoctor(mockcmd.MockExecFailure)
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
