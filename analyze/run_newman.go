package analyze

import (
	"os/exec"
)

// CommandRunner defines the interface for running commands.
type CommandRunner interface {
	Run() error
}

// RealCommand implements the CommandRunner interface using exec.Cmd.
type RealCommand struct {
	cmd *exec.Cmd
}

// Run executes the command.
func (c *RealCommand) Run() error {
	return c.cmd.Run()
}

// RunNewman runs Newman with specified collection and report paths.
func RunNewman(collectionFile string, reportPath string) error {
	cmd := exec.Command("newman", "run", collectionFile, "--reporters", "json", "--reporter-json-export", reportPath)
	return runCmd(cmd)
}

// Function to run command, can be overridden for testing
var runCmd = func(cmd *exec.Cmd) error {
	return cmd.Run()
}
