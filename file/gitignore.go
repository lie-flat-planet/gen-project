package file

import "strings"

type GitIgnore struct{}

func (g *GitIgnore) Name() string {
	return ".gitignore"
}

func (g *GitIgnore) Content() string {
	return strings.TrimSpace(`
.idea
.vscode

local.yml
.cursorignore
.cursorrules
`)
}
