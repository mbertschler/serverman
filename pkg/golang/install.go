package golang

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mbertschler/serverman/archive"
	"github.com/mbertschler/serverman/download"
)

const (
	goInstallPath = "/usr/local"
	goDownloadURL = "https://dl.google.com/go/%s.linux-amd64.tar.gz"
)

var (
	goDownloadDir = filepath.Join(os.TempDir(), "bootstrap", "go")
)

// Uninstall removes an existing Go installation.
// It does nothing if Go is not installed.
func Uninstall() error {
	return os.RemoveAll(goInstallPath)
}

// Install downloads the newest Go version and installs it on the system.
func Install() error {
	// download the newest version first
	version, err := NewestVersion()
	if err != nil {
		return err
	}
	url := fmt.Sprintf(goDownloadURL, version)

	log.Println("downloading", version)
	file, err := download.Download(url, goDownloadDir)
	if err != nil {
		return err
	}
	defer os.RemoveAll(goDownloadDir)

	log.Println("uninstalling old go version")
	// remove the old installation
	err = Uninstall()
	if err != nil {
		return err
	}

	log.Println("installing", version)
	// unpack the archive
	err = archive.UnpackFile(goInstallPath, file)
	if err != nil {
		return err
	}

	return nil
}
