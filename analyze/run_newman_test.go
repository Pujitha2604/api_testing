package analyze
 
import (
    "fmt"
    "os/exec"
    "testing"
 
    "github.com/stretchr/testify/mock"
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
    args := m.Called(cmd)
    return args.Error(0)
}
func RunNewmanWithRunner(collectionFile string, reportPath string, runner CmdRunner) error {
    cmd := exec.Command("newman", "run", collectionFile, "--reporters", "json", "--reporter-json-export", reportPath)
    err := runner.Run(cmd)
    if err != nil {
        return fmt.Errorf("error running Newman: %v", err)
    }
    return nil
}
 
func TestRunNewman(t *testing.T) {
    // Mock the CmdRunner for testing
    mockRunner := new(MockCmdRunner)
 
    // Define the expected command to be executed by RunNewman
    expectedCmd := exec.Command("newman", "run", "test_collection.json", "--reporters", "json", "--reporter-json-export", "test_report.json")
 
    // Mock the Run method of CmdRunner
    mockRunner.On("Run", expectedCmd).Return(nil) // Mock the command to return nil, indicating success
 
    // Call the function under test, passing the mock runner
    err := RunNewmanWithRunner("test_collection.json", "test_report.json", mockRunner)
 
    // Assert the result
    if err != nil {
        t.Errorf("Expected no error, got error running Newman: %v", err)
    }
 
    // Assert that the mock runner's Run method was called once with the expected command
    mockRunner.AssertCalled(t, "Run", expectedCmd)
}
 