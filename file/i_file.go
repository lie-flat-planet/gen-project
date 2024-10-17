package file

type IFile interface {
	Name() string
	Content() string
}
