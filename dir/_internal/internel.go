package _internal

import (
	"github.com/lie-flat-planet/gen-project/dir/_internal/controller"
	"github.com/lie-flat-planet/gen-project/dir/_internal/model"
	"github.com/lie-flat-planet/gen-project/dir/_internal/request"
	"github.com/lie-flat-planet/gen-project/dir/_internal/response"
	"github.com/lie-flat-planet/gen-project/dir/_internal/serializer"
	"github.com/lie-flat-planet/gen-project/dir/_internal/service"
	"github.com/lie-flat-planet/gen-project/file"
	_interface "github.com/lie-flat-planet/gen-project/generator/interface"
)

type Internal struct {
}

func (i *Internal) GetName() string {
	return "internal"
}

func (i *Internal) Children() []_interface.IDir {
	return []_interface.IDir{
		&model.Model{},
		&service.Service{},
		&controller.Controller{},
		&request.Request{},
		&response.Response{},
		&serializer.Serializer{},
	}
}

func (i *Internal) Files() []file.IFile {
	return nil
}

func (i *Internal) Create(basePath string) error {
	return _interface.Create(basePath, i)
}
