package golang

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDownload(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	err := download()
	require.NoError(t, err)

	f, err := os.Stat(tmpDownloadFile)
	require.NoError(t, err)
	assert.Greater(t, f.Size(), int64(1000000))

	err = remove()
	require.NoError(t, err)
}
