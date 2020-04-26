// Package test consists of helper for testing serverman.
package test

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"testing"

	"github.com/mbertschler/serverman"
	"github.com/mbertschler/serverman/pkg/sh"
)

var (
	dockerAvailableChecked bool
	dockerAvailable        bool

	runningContainersLock sync.Mutex
	runningContainers     = map[string]struct{}{}
)

func skipIfDockerUnavailable(t *testing.T) {
	if !dockerAvailableChecked {
		_, err := sh.RunString("docker", "version")
		if err == nil {
			dockerAvailable = true
		}
		dockerAvailableChecked = true
	}
	if !dockerAvailable {
		t.Skip("skipping because Docker is unavailable")
	}
}

// StartDebian starts a Debian Docker container and returns an Env that
// can run commands inside it. After use the cleanup function must be
// called, ideally in a defer right after this call.
func StartDebian(t *testing.T) (e *serverman.Env, cleanup func()) {
	image := "debian:10"
	skipIfDockerUnavailable(t)
	out, err := sh.RunString("docker", "images", "-q", image)
	if err != nil {
		t.Fatal("failed to check Debian Docker image:", err)
		return
	}
	if out == "" {
		out, err := sh.RunString("docker", "pull", image)
		if err != nil {
			t.Fatal("failed to pull Debian Docker image:", err, out)
			return
		}
	}
	id, err := sh.RunString("docker", "run", "-d", image, "sleep", "1m")
	if err != nil {
		t.Fatal("failed to start Debian Docker container:", err)
		return
	}
	id = strings.TrimSpace(id)
	runningContainersLock.Lock()
	runningContainers[id] = struct{}{}
	runningContainersLock.Unlock()
	cleanup = func() {
		_, err := sh.RunString("docker", "rm", "-f", id)
		if err != nil {
			log.Println(err)
		}
		runningContainersLock.Lock()
		delete(runningContainers, id)
		runningContainersLock.Unlock()
	}
	env := &serverman.Env{
		TestDockerID: id,
	}
	return env, cleanup
}

// StopAllContainersOnInterrupt waits for a SIGINT, stops all containers and
// calls os.Exit(1).
func StopAllContainersOnInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Println("interrupt, cleaning up all containers")
		StopAllContainers()
		os.Exit(1)
	}()
}

// StopAllContainers stops all Docker containers that were started by this package.
func StopAllContainers() {
	runningContainersLock.Lock()
	defer runningContainersLock.Unlock()
	for id := range runningContainers {
		_, err := sh.RunString("docker", "rm", "-f", id)
		if err != nil {
			log.Println(err)
		}
	}
}
