package apt

import (
	"os"
	"testing"

	"github.com/mbertschler/serverman/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	require.NoError(t, err)
	assert.False(t, ok)

	err = pkg.Apply(env)
	assert.NoError(t, err)
}

func TestAptInstallInvalid(t *testing.T) {
	env, stop := test.StartDebian(t)
	defer stop()

	pkg := Package{Name: "nanos"}
	ok, err := pkg.Check(env)
	assert.NoError(t, err)
	assert.False(t, ok)

	err = pkg.Apply(env)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "Unable to locate package nanos")
}
