package api

import (
	"github.com/lie-flat-planet/gen-project/dir/api/router"
	"github.com/lie-flat-planet/gen-project/file"
	_interface "github.com/lie-flat-planet/gen-project/generator/interface"
)

type API struct {
}

func (a *API) GetName() string {
	return "api"
}

func (a *API) Children() []_interface.IDir {
	return []_interface.IDir{
		&router.Router{},
	}
}

func (a *API) Files() []file.IFile {
	return nil
}

func (a *API) Create(basePath string) error {
	return _interface.Create(basePath, a)
}
