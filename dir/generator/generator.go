package Generator

import "github.com/lie-flat-planet/gen-project/global"

func Generate(basePath, moduleName string) error {
	global.SetModuleName(moduleName)

	r := newRoot(global.GetProjectName())
	return r.Create(basePath)
}
