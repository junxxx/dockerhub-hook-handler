package runner

import (
	"fmt"
	"os/exec"
	"syscall"
)

// run command and wait it's exit status code
func RunAndWait(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return fmt.Errorf("Exit Status: %d", status.ExitStatus())
			}
		} else {
			return fmt.Errorf("cmd.Wait: %v", err)
		}
	}
	return nil
}
