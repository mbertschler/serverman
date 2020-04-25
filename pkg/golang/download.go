package golang

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	goDownloadURL = "https://dl.google.com/go/%s.linux-amd64.tar.gz"
)

var (
	tmpDownloadDir  = filepath.Join(os.TempDir(), "bootstrap", "go")
	tmpDownloadFile = filepath.Join(tmpDownloadDir, "go.tar.gz")
)

func download() error {
	// get newest go version
	version, err := NewestVersion()
	if err != nil {
		return err
	}

	// download file
	url := fmt.Sprintf(goDownloadURL, version)
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// create output directory
	err = os.MkdirAll(tmpDownloadDir, 0755)
	if err != nil {
		return err
	}

	// create output file
	f, err := os.Create(tmpDownloadFile)
	if err != nil {
		return err
	}
	defer f.Close()

	// copy body to file
	_, err = io.Copy(f, res.Body)
	return err
}

func remove() error {
	return os.RemoveAll(tmpDownloadDir)
}
