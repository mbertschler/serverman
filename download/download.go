package download

import (
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// Download downloads a file from a URL, stores it into the given directory and returns the name of the written file.
// It fails if four attempts return an error.
func Download(url, dir string) (string, error) {
	var file string
	var err error
	for i := 0; i < 4; i++ {
		file, err = download(url, dir)
		if err == nil {
			break
		}
	}

	return file, err
}

func download(url, dir string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// create output directory
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return "", err
	}

	file := filepath.Join(dir, path.Base(url))

	// create output file
	f, err := os.Create(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// copy body to file
	_, err = io.Copy(f, res.Body)
	return file, err
}
