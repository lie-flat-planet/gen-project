package api

import "github.com/lie-flat-planet/gen-project/file"

type API struct {
}

func (a *API) Name() string {
	return "api"
}

func (a *API) Children() []IDir {
	return []IDir{
		&Route{},
	}
}

func (a *API) Files() []file.IFile {
	return nil
}

func (a *API) Create(basePath string) error {
	return create(basePath, a)
}
