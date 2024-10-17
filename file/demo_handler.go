package file

import (
	"fmt"
	"github.com/lie-flat-planet/gen-project/global"
	"strings"
)

type DemoHandler struct{}

func (*DemoHandler) Name() string {
	return "demo.go"
}

func (*DemoHandler) Content() string {
	moduleName := global.GetModuleName()

	return strings.TrimSpace(fmt.Sprintf(`
package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/lie-flat-planet/httputil"
	"%s/api/controller"
	"%s/config"
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
`, moduleName, moduleName))
}
