package dir

import "github.com/lie-flat-planet/gen-project/file"

type Route struct{}

func (r *Route) Name() string {
	return "route"
}

func (r *Route) ListChildren() []IDir {

}

func (r *Route) ListFile() []file.IFile {

}
