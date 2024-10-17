package config

import (
	"github.com/lie-flat-planet/gen-project/file"
	_interface "github.com/lie-flat-planet/gen-project/generator/interface"
)

type Config struct{}

func (*Config) GetName() string {
	return "config"
}

func (*Config) Children() []_interface.IDir {
	return nil
}

func (*Config) Files() []file.IFile {
	return []file.IFile{
		&file.Config{},
		&file.LocalYML{},
	}
}

func (config *Config) Create(basePath string) error {
	return _interface.Create(basePath, config)
}
