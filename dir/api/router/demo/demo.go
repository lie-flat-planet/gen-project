package demo

import (
	"github.com/lie-flat-planet/gen-project/file"
	_interface "github.com/lie-flat-planet/gen-project/generator/interface"
)

type Demo struct {
}

func (*Demo) GetName() string {
	return "demo"
}

func (*Demo) Children() []_interface.IDir {
	return nil
}

func (*Demo) Files() []file.IFile {
	return []file.IFile{
		&file.DemoHandler{},
		&file.DemoRouter{},
	}
}

func (demo *Demo) Create(basePath string) error {
	return _interface.Create(basePath, demo)
}
