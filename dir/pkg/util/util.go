package util

import (
	"github.com/lie-flat-planet/gen-project/file"
	_interface "github.com/lie-flat-planet/gen-project/generator/interface"
)

type Util struct {
}

func (*Util) GetName() string {
	return "util"
}

func (*Util) Children() []_interface.IDir {
	return nil
}

func (*Util) Files() []file.IFile {
	return []file.IFile{
		&file.UtilJson{},
	}
}

func (util *Util) Create(basePath string) error {
	return _interface.Create(basePath, util)
}
