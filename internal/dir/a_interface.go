package dir

type IDIR interface {
	RootPath() string
	Path() string
	File() string
	FileContent() string
	SetProjectName(string)
}
