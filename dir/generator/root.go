package Generator

import (
	"github.com/lie-flat-planet/gen-project/dir/api/route"
	_interface "github.com/lie-flat-planet/gen-project/dir/interface"
	"github.com/lie-flat-planet/gen-project/file"
)

type root struct {
	name string
}

func newRoot(name string) *root {
	return &root{
		name: name,
	}
}

func (root *root) GetName() string {
	return root.name
}

func (root *root) Children() []_interface.IDir {
	return []_interface.IDir{
		&route.Route{},
	}
}

func (root *root) Files() []file.IFile {
	return []file.IFile{
		&file.GitIgnore{},
		&file.Version{},
		&file.Dockerfile{},
		&file.GoMod{},
		&file.Makefile{},
	}
}

func (root *root) Create(basePath string) error {
	return _interface.Create(basePath, root)
}
