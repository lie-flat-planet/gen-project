package file

import (
	"fmt"
	"github.com/lie-flat-planet/gen-project/global"
	"strings"
)

type GoMod struct{}

func (g *GoMod) Name() string {
	return "go.mod"
}

func (g *GoMod) Content() string {
	return strings.TrimSpace(fmt.Sprintf(`
module %s

go 1.24.0

require (
	github.com/gin-gonic/gin v1.10.0
	github.com/lie-flat-planet/httputil v0.0.3
	github.com/lie-flat-planet/service-init-tool v1.0.1
	github.com/swaggo/files v1.0.1
	github.com/swaggo/gin-swagger v1.6.0
	github.com/swaggo/swag v1.16.3
	gorm.io/gorm v1.25.12
	github.com/go-redis/redis/v8 v8.11.5
)

`, global.GetModuleName()))
}
