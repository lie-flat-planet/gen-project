package serializer

import (
	"github.com/lie-flat-planet/gen-project/file"
	_interface "github.com/lie-flat-planet/gen-project/generator/interface"
)

type Serializer struct {
}

func (*Serializer) GetName() string {
	return "serializer"
}

func (*Serializer) Children() []_interface.IDir {
	return nil
}

func (*Serializer) Files() []file.IFile {
	return []file.IFile{}
}

func (ser *Serializer) Create(basePath string) error {
	return _interface.Create(basePath, ser)
}
