package apt

import (
	"os"
	"testing"

	"github.com/mbertschler/serverman/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	test.StopAllContainersOnInterrupt()
	os.Exit(m.Run())
}

func TestAptInstall(t *testing.T) {
	env, stop := test.StartDebian(t)
	defer stop()

	pkg := Package{Name: "nano"}
	ok, err := pkg.Check(env)
	assert.False(t, ok)
	assert.NoError(t, err)

	err = pkg.Apply(env)
	assert.NoError(t, err)
}

func TestAptInstallInvalid(t *testing.T) {
	env, stop := test.StartDebian(t)
	defer stop()

	pkg := Package{Name: "nanos"}
	ok, err := pkg.Check(env)
	assert.False(t, ok)
	assert.NoError(t, err)

	err = pkg.Apply(env)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Unable to locate package nanos")
}
