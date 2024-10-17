package docs

import (
	"github.com/lie-flat-planet/gen-project/file"
	_interface "github.com/lie-flat-planet/gen-project/generator/interface"
)

type Docs struct {
}

func (*Docs) GetName() string {
	return "docs"
}

func (*Docs) Children() []_interface.IDir {
	return nil
}

func (*Docs) Files() []file.IFile {
	return []file.IFile{
		&file.Docs{},
	}
}

func (docs *Docs) Create(basePath string) error {
	return _interface.Create(basePath, docs)
}
