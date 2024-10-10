package dir

import "github.com/lie-flat-planet/gen-project/file"

type IDir interface {
	Name() string
	ListChildren() []IDir
	ListFile() []file.IFile
}
