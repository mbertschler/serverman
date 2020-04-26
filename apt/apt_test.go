package apt

import (
	"os"
	"testing"
	"time"

	"github.com/mbertschler/serverman/pkg/test"
)

func TestMain(m *testing.M) {
	test.StopAllContainersOnInterrupt()
	os.Exit(m.Run())
}

func TestAptInstall(t *testing.T) {
	stop := test.StartDebian(t)
	defer stop()
	t.Log("start worked... sleeping")
	time.Sleep(10 * time.Minute)
}
