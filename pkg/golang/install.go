package golang

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
)

const (
	goInstallPath = "/usr/local"
)

// Uninstall removes an existing Go installation.
// It does nothing if Go is not installed.
func Uninstall() error {
	return os.RemoveAll(goInstallPath)
}

// Install downloads the newest Go version and installs it on the system.
func Install() error {
	// download the newest version first
	version, _ := NewestVersion()
	log.Println("downloading", version)
	err := download()
	if err != nil {
		return err
	}
	// defer remove()
	log.Println("done")

	// remove the old installation
	// err = Uninstall()
	// if err != nil {
	// 	return err
	// }

	// unpack the archive
	err = unpack()
	if err != nil {
		return err
	}

	return nil
}

func unpack() error {
	// get the downloaded file
	dlFile, err := os.Open(tmpDownloadFile)
	if err != nil {
		return err
	}
	defer dlFile.Close()

	// un-gzip it
	gz, err := gzip.NewReader(dlFile)
	if err != nil {
		return err
	}

	// read the tar
	tr := tar.NewReader(gz)

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

		target := filepath.Join(goInstallPath, header.Name)

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
