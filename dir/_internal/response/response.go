package response

import (
	"github.com/lie-flat-planet/gen-project/file"
	_interface "github.com/lie-flat-planet/gen-project/generator/interface"
)

type Response struct {
}

func (*Response) GetName() string {
	return "response"
}

func (*Response) Children() []_interface.IDir {
	return nil
}

func (*Response) Files() []file.IFile {
	return nil
}

func (resp *Response) Create(basePath string) error {
	return _interface.Create(basePath, resp)
}
