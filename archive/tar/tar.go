package tar

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Unpack extracts a tar.
func Unpack(r io.Reader, dest string) error {
	// read the tar
	tr := tar.NewReader(r)

	for {
		// switch to next object
		header, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				// done
				return nil
			}
			return err
		}

		log.Println(header.Name)

		target := filepath.Join(dest, header.Name)

		switch header.Typeflag {
		// directory
		case tar.TypeDir:
			err = os.MkdirAll(target, 0755)
			if err != nil {
				return err
			}

			// file
		case tar.TypeReg:
			err = writeFile(target, tr, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
		}
	}
}

func writeFile(path string, r io.Reader, perm os.FileMode) error {
	// open output file
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, perm)
	if err != nil {
		return err
	}
	defer f.Close()

	// write to file
	_, err = io.Copy(f, r)
	return err
}
