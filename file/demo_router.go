package file

import "strings"

type DemoRouter struct{}

func (*DemoRouter) Name() string {
	return "z_router.go"
}

func (*DemoRouter) Content() string {
	return strings.TrimSpace(`
package demo

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r = r.Group("demo")

	r.GET("/hello-world", FetchFirst)
}


`)
}