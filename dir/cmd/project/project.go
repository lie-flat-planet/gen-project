package project

import (
	"github.com/lie-flat-planet/gen-project/dir/cmd/project/docs"
	"github.com/lie-flat-planet/gen-project/file"
	_interface "github.com/lie-flat-planet/gen-project/generator/interface"
)

type Project struct {
	Name string
}

func (project *Project) GetName() string {
	if project.Name == "" {
		panic("project name can't be empty")
	}

	return project.Name
}

func (*Project) Children() []_interface.IDir {
	return []_interface.IDir{
		&docs.Docs{},
	}
}

func (*Project) Files() []file.IFile {
	return []file.IFile{
		&file.Main{},
	}
}

func (project *Project) Create(basePath string) error {
	return _interface.Create(basePath, project)
}
