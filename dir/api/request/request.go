package request

import (
	"github.com/lie-flat-planet/gen-project/file"
	_interface "github.com/lie-flat-planet/gen-project/generator/interface"
)

type Request struct {
}

func (*Request) GetName() string {
	return "request"
}

func (*Request) Children() []_interface.IDir {
	return nil
}

func (*Request) Files() []file.IFile {
	return nil
}

func (request *Request) Create(basePath string) error {
	return _interface.Create(basePath, request)
}
