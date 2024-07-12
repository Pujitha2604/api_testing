package analyze

import (
	"fmt"
	"os/exec"
)

func RunNewman(collectionFile string, reportPath string) error {
	cmd := exec.Command("newman", "run", collectionFile, "--reporters", "json", "--reporter-json-export", reportPath)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error running Newman: %v", err)
	}
	return err
}
