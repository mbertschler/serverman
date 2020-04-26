package apt

import "errors"

type Package struct {
	Name string
}

func (p *Package) Apply() (changed bool, err error) {
	return false, errors.New("not implemented")
}
