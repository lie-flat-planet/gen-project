package file

import (
	"fmt"
	"github.com/lie-flat-planet/gen-project/global"
	"strings"
)

type Main struct {
}

func (*Main) Name() string {
	return "main.go"
}

func (*Main) Content() string {
	return strings.TrimSpace(fmt.Sprintf(`
package main

import (
	"%s/api/router"
)

func main() {
	router.Start()
}

`, global.GetModuleName()))
}
