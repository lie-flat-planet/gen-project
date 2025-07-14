package controller

import (
	"github.com/lie-flat-planet/gen-project/file"
	_interface "github.com/lie-flat-planet/gen-project/generator/interface"
)

type Controller struct {
}

func (*Controller) GetName() string {
	return "controller"
}

func (*Controller) Children() []_interface.IDir {
	return nil
}

func (*Controller) Files() []file.IFile {
	return []file.IFile{
		&file.ControllerDemo{},
	}
}

func (ctl *Controller) Create(basePath string) error {
	return _interface.Create(basePath, ctl)
}
