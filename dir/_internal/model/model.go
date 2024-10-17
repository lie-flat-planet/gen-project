package model

import (
	"github.com/lie-flat-planet/gen-project/file"
	_interface "github.com/lie-flat-planet/gen-project/generator/interface"
)

type Model struct{}

func (m *Model) GetName() string {
	return "model"
}

func (m *Model) Children() []_interface.IDir {
	return nil
}

func (m *Model) Files() []file.IFile {
	return []file.IFile{
		&file.ModelDb{},
		&file.ModelUser{},
	}
}

func (m *Model) Create(basePath string) error {
	return _interface.Create(basePath, m)
}
