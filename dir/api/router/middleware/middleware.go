package middleware

import (
	"github.com/lie-flat-planet/gen-project/file"
	_interface "github.com/lie-flat-planet/gen-project/generator/interface"
)

type Middleware struct {
}

func (*Middleware) GetName() string {
	return "middleware"
}

func (*Middleware) Children() []_interface.IDir {
	return nil
}

func (*Middleware) Files() []file.IFile {
	return nil
}

func (middleware *Middleware) Create(basePath string) error {
	return _interface.Create(basePath, middleware)
}
