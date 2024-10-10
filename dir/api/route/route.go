package route

import (
	_interface "github.com/lie-flat-planet/gen-project/dir/interface"
	"github.com/lie-flat-planet/gen-project/file"
)

type Route struct {
}

func (r *Route) Name() string {
	return "route"
}

func (r *Route) Children() []_interface.IDir {
	return nil
}

func (r *Route) Files() []file.IFile {
	return []file.IFile{
		&file.ZRoute{},
	}
}

func (r *Route) Create(basePath string) error {
	return _interface.Create(basePath, r)
}
