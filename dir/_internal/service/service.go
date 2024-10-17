package service

import (
	"github.com/lie-flat-planet/gen-project/file"
	_interface "github.com/lie-flat-planet/gen-project/generator/interface"
)

type Service struct {
}

func (s *Service) GetName() string {
	return "service"
}

func (s *Service) Children() []_interface.IDir {
	return nil
}

func (s *Service) Files() []file.IFile {
	return []file.IFile{
		&file.ServiceDemo{},
	}
}

func (s *Service) Create(basePath string) error {
	return _interface.Create(basePath, s)
}
