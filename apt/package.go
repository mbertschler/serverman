// Package apt installs Debian packages with apt-get.
package apt

import (
	"strings"

	"github.com/mbertschler/serverman"
	"github.com/pkg/errors"
)

// Package represents an apt package that should be installed.
type Package struct {
	Name string
}

// Check checks if the apt package is already installed using dpkg.
func (p *Package) Check(e *serverman.Env) (ok bool, err error) {
	out, err := e.RunString("dpkg", "-l", p.Name)
	if err == nil {
		return true, nil
	}
	if strings.Contains(out, "no packages found matching") {
		return false, nil
	}
	return false, errors.Wrapf(err, "unexpected output %q", out)
}

// Apply installs the package with apt-get.
func (p *Package) Apply(e *serverman.Env) (err error) {
	out, err := e.RunString("apt-get", "update")
	if err != nil {
		return errors.Wrapf(err, "apt-get update failed with: %q", out)
	}
	out, err = e.RunString("apt-get", "install", "-yq", p.Name)
	if err != nil {
		return errors.Wrapf(err, "apt-get install failed with: %q", out)
	}
	return nil
}

// Remove purges the package with apt-get.
func (p *Package) Remove(e *serverman.Env) (err error) {
	out, err := e.RunString("apt-get", "purge", "-y", p.Name)
	if err != nil {
		return errors.Wrapf(err, "apt-get purge failed with: %q", out)
	}
	return nil
}
