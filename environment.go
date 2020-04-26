// Package serverman is the top level framework package of serverman.
package serverman

import (
	"github.com/mbertschler/serverman/pkg/sh"
)

// Env represents the environment in which the serverman config is applied.
// For testing it can "docker exec" the command if the TestDockerID field is set.
type Env struct {
	TestDockerID string
}

// RunString executes the passed command in this environment.
func (e *Env) RunString(name string, args ...string) (string, error) {
	if e.TestDockerID != "" {
		args = append([]string{"exec", "-t", e.TestDockerID, name}, args...)
		name = "docker"
	}
	return sh.RunString(name, args...)
}
