package dir

import "fmt"

func init() {
	Dirs = append(Dirs, NewGitIgnore())
}

type GitIgnore struct {
	projectName string
}

func NewGitIgnore() *GitIgnore {
	return new(GitIgnore)
}

func (g *GitIgnore) SubPath() string {
	return ""
}

func (g *GitIgnore) Path() string {
	return ""
}

func (g *GitIgnore) File() string {
	return ".gitignore"
}

func (g *GitIgnore) FileContent() string {
	return fmt.Sprintf(`.idea/
.vscode/
dockerfile.default.yml
cmd/%s/%s
`, g.projectName, g.projectName)
}

func (g *GitIgnore) SetProjectName(name string) {
	g.projectName = name
}
