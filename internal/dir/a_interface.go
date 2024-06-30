package dir

type IDIR interface {
	SubPath() string
	Path() string
	File() string
	FileContent() string
	SetProjectName(string)
}
