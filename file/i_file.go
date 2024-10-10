package file

// File represents a file with a name and content
type File struct {
	Name    string // File name
	Content string // File content
}

type IFile interface {
	Name() string
	Content() string
}
