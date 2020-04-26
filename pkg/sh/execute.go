package sh

import (
	"bytes"
	"os/exec"
)

// RunString executes the command with arguments and returns the output
// string of stderr and stdout and an error if the command fails.
func RunString(name string, args ...string) (string, error) {
	buf := &bytes.Buffer{}
	cmd := exec.Command(name, args...)
	cmd.Stderr = buf
	cmd.Stdout = buf
	err := cmd.Run()
	return buf.String(), err
}
