package pkg

import (
	"os"
	"os/exec"
)

//Execute runs a command
func Execute(command string, args []string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}
