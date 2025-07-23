package generator

import (
	"github.com/lie-flat-planet/gen-project/dir/_internal"
	"github.com/lie-flat-planet/gen-project/dir/api"
	"github.com/lie-flat-planet/gen-project/dir/cmd"
	"github.com/lie-flat-planet/gen-project/dir/config"
	"github.com/lie-flat-planet/gen-project/dir/pkg"
	"github.com/lie-flat-planet/gen-project/file"
	_interface "github.com/lie-flat-planet/gen-project/generator/interface"
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
		&api.API{},
		&cmd.CMD{},
		&config.Config{},
		&_internal.Internal{},
		&pkg.PKG{},
	}
}

func (root *root) Files() []file.IFile {
	return []file.IFile{
		&file.CursorIgnore{},
		&file.CursorRules{},
		&file.GitIgnore{},
		&file.Version{},
		&file.Dockerfile{},
		&file.DockerfileDev{},
		&file.GoMod{},
		&file.Makefile{},
	}
}

func (root *root) Create(basePath string) error {
	return _interface.Create(basePath, root)
}
