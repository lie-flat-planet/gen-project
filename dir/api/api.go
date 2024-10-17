package api

import (
	_interface "github.com/lie-flat-planet/gen-project/dir/interface"
	"github.com/lie-flat-planet/gen-project/file"
)

type API struct {
}

func (a *API) Name() string {
	return "api"
}

func (a *API) Children() []_interface.IDir {
	return []_interface.IDir{
		&Route{},
	}
}

func (a *API) Files() []file.IFile {
	return nil
}

func (a *API) Create(basePath string) error {
	return create(basePath, a)
}
