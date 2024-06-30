package dir

import "fmt"

func init() {
	Dirs = append(Dirs, NewMod())
}

type Mod struct {
	projectName string
}

func NewMod() *Mod {
	return new(Mod)
}

func (m *Mod) SubPath() string {
	return ""
}

func (m *Mod) Path() string {
	return ""
}

func (m *Mod) File() string {
	return "go.mod"
}

func (m *Mod) FileContent() string {

	return fmt.Sprintf(`// TODO: replace xxx to group
module github.com/xxx/%s
// TODO: replace go version
go 1.20

require (
	github.com/go-courier/statuserror v1.2.1
	github.com/gin-gonic/gin v1.10.0
	github.com/kunlun-qilian/conflogger v0.2.0
	github.com/kunlun-qilian/confpostgres v0.0.2
	github.com/kunlun-qilian/confserver v0.14.4
	github.com/lie-flat-planet/confx v1.0.1
	github.com/spf13/cobra v1.8.0
)`, m.projectName)
}

func (m *Mod) SetProjectName(name string) {
	m.projectName = name
}
