package _interface

import "github.com/lie-flat-planet/gen-project/file"

type IDir interface {
	GetName() string
	Children() []IDir
	Files() []file.IFile
	Create(path string) error
}
