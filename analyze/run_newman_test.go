package analyze

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/mock"
)

var (
	runner CmdRunner = &RealCmdRunner{} // Default implementation using real command runner
)

// CmdRunner interface represents a command runner (for mocking purposes)
type CmdRunner interface {
	Run(*exec.Cmd) error
}

// RealCmdRunner is a real implementation of CmdRunner using exec.Cmd
type RealCmdRunner struct{}

// Run implements CmdRunner interface for executing a real command
func (r *RealCmdRunner) Run(cmd *exec.Cmd) error {
	return cmd.Run()
}

// MockCmdRunner is a mock implementation of CmdRunner for testing
type MockCmdRunner struct {
	mock.Mock
}

// Run implements CmdRunner interface for mock command execution
func (m *MockCmdRunner) Run(cmd *exec.Cmd) error {
	m.Called(cmd)
	return nil // Simulate successful execution
}
func TestRunNewman(t *testing.T) {
	// Mock the CmdRunner for testing
	mockRunner := new(MockCmdRunner)

	// Replace the package-level runner variable with the mock instance
	runner = mockRunner

	// Define the expected command to be executed by RunNewman
	expectedCmd := exec.Command("C:\\Users\\Rekanto\\AppData\\Roaming\\npm\\newman.cmd", "run", "test_collection.json", "--reporters", "json", "--reporter-json-export", "test_report.json")

	// Mock the Run method of CmdRunner
	mockRunner.On("Run", expectedCmd).Return(nil) // Expect Run to be called with the expected command

	// Call the function under test
	err := RunNewman("test_collection.json", "test_report.json")

	// Assert the result
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Assert that the mock runner's Run method was called once with the expected command
	mockRunner.AssertCalled(t, "Run", expectedCmd)
}


