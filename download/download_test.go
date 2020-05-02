package download

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDownload(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	dir := filepath.Join(os.TempDir(), fmt.Sprint(time.Now().UnixNano()))
	file, err := Download("https://www.exahome.net/img/exahome-logo.svg", dir)
	require.NoError(t, err)

	f, err := os.Stat(file)
	require.NoError(t, err)
	assert.Greater(t, f.Size(), int64(1e3))

	err = os.RemoveAll(dir)
	require.NoError(t, err)
}
