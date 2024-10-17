package global

import "strings"

var moduleName string

func SetModuleName(name string) {
	moduleName = name
}

func GetModuleName() string {
	return moduleName
}

func GetProjectName() string {
	return parseProjectName(moduleName)
}

func parseProjectName(name string) string {
	seg := strings.Split(name, "/")

	segLen := len(seg)

	return seg[segLen-1]
}
