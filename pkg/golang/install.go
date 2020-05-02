package golang

import (
	"log"
	"os"

	"github.com/mbertschler/serverman/archive"
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
	defer remove()

	log.Println("uninstalling old go version")
	// remove the old installation
	err = Uninstall()
	if err != nil {
		return err
	}

	log.Println("installing", version)
	// unpack the archive
	err = unpack()
	if err != nil {
		return err
	}

	return nil
}

func unpack() error {
	return archive.UnpackFile(goInstallPath, tmpDownloadFile)
}
