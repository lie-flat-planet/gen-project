package router

import (
	"github.com/lie-flat-planet/gen-project/dir/api/router/demo"
	"github.com/lie-flat-planet/gen-project/dir/api/router/middleware"
	"github.com/lie-flat-planet/gen-project/file"
	_interface "github.com/lie-flat-planet/gen-project/generator/interface"
)

type Router struct {
	Name string
}

func (r *Router) GetName() string {
	if r.Name == "" {
		return "router"
	}
	return r.Name
}

func (r *Router) Children() []_interface.IDir {
	return []_interface.IDir{
		&demo.Demo{},
		&middleware.Middleware{},
	}
}

func (r *Router) Files() []file.IFile {
	return []file.IFile{
		&file.Route{},
	}
}

func (r *Router) Create(basePath string) error {
	return _interface.Create(basePath, r)
}
