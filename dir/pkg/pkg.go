package pkg

import (
	"github.com/lie-flat-planet/gen-project/dir/pkg/util"
	"github.com/lie-flat-planet/gen-project/file"
	_interface "github.com/lie-flat-planet/gen-project/generator/interface"
)

type PKG struct {
}

func (*PKG) GetName() string {
	return "pkg"
}

func (*PKG) Children() []_interface.IDir {
	return []_interface.IDir{
		&util.Util{},
	}
}

func (*PKG) Files() []file.IFile {
	return nil
}

func (pkg *PKG) Create(basePath string) error {
	return _interface.Create(basePath, pkg)
}
