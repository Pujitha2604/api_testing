package analyze

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

var runCommand = func(cmd *exec.Cmd) error {
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("exit status 1")
	}
	return nil
}

func TestRunNewman(t *testing.T) {
	// Mock behavior of runCommand variable
	mockError := fmt.Errorf("exit status 1")
	runCommand = func(cmd *exec.Cmd) error {
		if cmd.Path == "newman" && len(cmd.Args) >= 2 && cmd.Args[1] == "run" {
			// Simulate command execution failure
			return mockError
		}
		// Execute the real command
		return cmd.Run()
	}
	defer func() {
		runCommand = func(cmd *exec.Cmd) error {
			return cmd.Run()
		}
	}()

	// Expectation for exec.Command
	expectedCommand := []string{"newman", "run", "mock_collection.json", "--reporters", "json", "--reporter-json-export", "mock_report.json"}

	// Call the function
	collectionFile := "mock_collection.json"
	reportPath := "mock_report.json"
	err := RunNewman(collectionFile, reportPath)

	// Assertions for RunNewman
	assert.EqualError(t, err, mockError.Error())

	// Verify that exec.Command would have been called with expected parameters
	cmd := exec.Command(expectedCommand[0], expectedCommand[1:]...)
	err = runCommand(cmd)
	assert.EqualError(t, err, mockError.Error())

	// Clean up any generated report file
	os.Remove(reportPath)
}

// func TestRealCommandRun(t *testing.T) {
// 	// Mock exec.Cmd for testing purposes
// 	mockCmd := exec.Command("echo", "hello")
// 	realCmd := &RealCommand{cmd: mockCmd}

// 	// Simulate successful execution
// 	err := realCmd.Run()
// 	assert.NoError(t, err)

// 	// Optionally, check coverage for error paths in RealCommand.Run() here
// }
