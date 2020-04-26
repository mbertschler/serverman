// Package sh is a helper for running commands with os/exec
package sh

import (
	"os/exec"
)

// RunString executes the command with arguments and returns the output
// string of stderr and stdout and an error if the command fails.
func RunString(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	buf, err := cmd.CombinedOutput()
	return string(buf), err
}
