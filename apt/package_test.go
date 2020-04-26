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

func TestAptInstallRemove(t *testing.T) {
	env, stop := test.StartDebian(t)
	defer stop()

	pkg := Package{Name: "nano"}

	// not yet installed > check should be false
	ok, err := pkg.Check(env)
	require.NoError(t, err)
	assert.False(t, ok)

	err = pkg.Apply(env)
	assert.NoError(t, err)

	// installed > check should be true
	ok, err = pkg.Check(env)
	require.NoError(t, err)
	assert.True(t, ok)

	err = pkg.Remove(env)
	assert.NoError(t, err)

	// removed > check should be false again
	ok, err = pkg.Check(env)
	require.NoError(t, err)
	assert.False(t, ok)
}

func TestAptInvalidPackage(t *testing.T) {
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
