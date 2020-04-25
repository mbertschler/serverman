package golang

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewestVersion(t *testing.T) {
	v, err := NewestVersion()
	require.NoError(t, err)
	require.Contains(t, v, "go1")
}
