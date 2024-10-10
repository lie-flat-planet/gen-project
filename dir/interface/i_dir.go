package _interface

import "github.com/lie-flat-planet/gen-project/file"

type IDir interface {
	Name() string
	Children() []IDir
	Files() []file.IFile
	Create(path string) error
}
