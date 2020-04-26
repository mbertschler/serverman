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

func SkipIfDockerUnavailable(t *testing.T) {
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

func StartDebian(t *testing.T) (e *serverman.Env, cleanup func()) {
	SkipIfDockerUnavailable(t)
	id, err := sh.RunString("docker", "run", "-d", "debian:10", "sleep", "1m")
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
