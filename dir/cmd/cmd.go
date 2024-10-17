package cmd

import (
	"github.com/lie-flat-planet/gen-project/dir/cmd/project"
	"github.com/lie-flat-planet/gen-project/file"
	_interface "github.com/lie-flat-planet/gen-project/generator/interface"
	"github.com/lie-flat-planet/gen-project/global"
)

type CMD struct {
}

func (*CMD) GetName() string {
	return "cmd"
}

func (*CMD) Children() []_interface.IDir {
	return []_interface.IDir{
		&project.Project{
			Name: global.GetProjectName(),
		},
	}
}

func (*CMD) Files() []file.IFile {
	return nil
}

func (cmd *CMD) Create(basePath string) error {
	return _interface.Create(basePath, cmd)
}
