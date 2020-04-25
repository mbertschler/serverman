package golang

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	goVersionURL = "https://golang.org/VERSION?m=text"
)

// NewestVersion retrieves the version string of the newest go version.
func NewestVersion() (string, error) {
	res, err := http.Get(goVersionURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	version := string(buf)

	if !strings.HasPrefix(version, "go1") {
		return "", fmt.Errorf("invalid version result: %s", version)
	}

	return version, nil
}
