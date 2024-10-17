package file

import "strings"

type DemoHandler struct{}

func (*DemoHandler) Name() string {
	return "demo.go"
}

func (*DemoHandler) Content() string {
	return strings.TrimSpace(`
package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/lie-flat-planet/httputil"
	"test-gen-project/api/controller"
	"test-gen-project/config"
)

func FetchFirst(ctx *gin.Context) {
	ctl := controller.FactoryDemo()

	user, err := ctl.FetchFirst(ctx)

	(&httputil.RESP{
		Content:     user,
		ServiceCode: config.Config.Server.Code,
		Err:         err,
	}).Output(ctx)

	return
}
`)
}