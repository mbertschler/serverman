package archive

import (
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/mbertschler/serverman/archive/tar"
)

// UnpackFile unpacks any supported archive to the destination directory.
// If the destination directory doesn't exist, it will be created.
func UnpackFile(dst, src string) error {
	// create destination directory if it doesn't exist
	err := os.MkdirAll(dst, 0755)
	if err != nil {
		return err
	}

	// open archive file
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	r := io.Reader(f)

	// get extension and inflate gz
	ext := filepath.Ext(src)
	if ext == ".gz" {
		src = strings.TrimSuffix(src, ".gz")
		ext = filepath.Ext(src)

		r, err = gzip.NewReader(r)
	}

	switch ext {
	case ".tar":
		return tar.Unpack(r, dst)
	}
	return nil
}
