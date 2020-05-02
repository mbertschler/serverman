package gz

import (
	"compress/gzip"
	"io"
)

func unpack() (io.Reader, error) {
	// un-gzip it
	gz, err := gzip.NewReader(dlFile)
	if err != nil {
		return err
	}
}
