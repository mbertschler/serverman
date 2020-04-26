package serverman

import (
	"github.com/mbertschler/serverman/pkg/sh"
)

type Env struct {
	TestDockerID string
}

func (e *Env) RunString(name string, args ...string) (string, error) {
	if e.TestDockerID != "" {
		args = append([]string{"exec", "-t", e.TestDockerID, name}, args...)
		name = "docker"
	}
	return sh.RunString(name, args...)
}
