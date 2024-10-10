package Generator

import (
	"github.com/lie-flat-planet/gen-project/dir/api/route"
	_interface "github.com/lie-flat-planet/gen-project/dir/interface"
	"github.com/lie-flat-planet/gen-project/file"
)

func Generate(basePath, name string) error {
	r := newRoot(name)
	return r.Create(basePath)
}

type root struct {
	name string
}

func newRoot(name string) *root {
	return &root{
		name: name,
	}
}

func (root *root) Name() string {
	return root.name
}

func (root *root) Children() []_interface.IDir {
	return []_interface.IDir{
		&route.Route{},
	}
}

func (root *root) Files() []file.IFile {
	return nil
}

func (root *root) Create(basePath string) error {
	return _interface.Create(basePath, root)
}
