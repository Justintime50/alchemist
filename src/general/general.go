package general

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/justintime50/mockcmd/mockcmd"
	"github.com/natefinch/lumberjack"
)

// RunCommand runs a shell command
func RunCommand(cmdContext mockcmd.ExecContext, command string, args []string) (*bytes.Buffer, error) {
	cmd := cmdContext(command, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprint(err) + ": " + stderr.String())
	}

	return &out, nil
}

// SetupDir gets the user's home dir and appends the alchemist dir to it
func SetupDir(dir string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	alchemistUpdateDir := homeDir + "/alchemist/" + dir
	_ = os.MkdirAll(alchemistUpdateDir, os.ModePerm)

	return alchemistUpdateDir
}

// SetupLogging sets up logging
func SetupLogging(dir string, action string) string {
	logPath := dir + "/alchemist-" + action + ".log"
	log.SetOutput(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    1, // megabytes
		MaxBackups: 5,
		MaxAge:     90, // days
	})

	return logPath
}
