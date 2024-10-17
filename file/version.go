package file

import "strings"

type Version struct{}

func (v *Version) Name() string {
	return "version.go"
}

func (v *Version) Content() string {
	return strings.TrimSpace(`
1.0.0
`)
}
