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
`+"type UserBody struct {\n\tName   string  `json:\"name\"`\n\tSalary float64 `json:\"salary\"`\n\tAge    int     `json:\"age\"`\n}\n\t"+`
// FetchFirst
// @BasePath /
// PingExample godoc
// @Summary FetchFirst
// @Schemes
// @Description 查询用户第一条数据
// @Tags FetchFirst
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization bearer token"
// @Param Body body UserBody true "body"
// @Param userID query string true "用户id"
// @Success 200 {object} resp.UserRESP 成功
// @Failure 400 {object} httputil.ErrorRESP 失败
// @Failure 500 {object} httputil.ErrorRESP 失败
// @Router /monitor-gateway/api/v1/demo/hello-world [POST]
// @ID FetchFirst
func FetchFirst(ctx *gin.Context) {
	userID := ctx.Query("userID")
	fmt.Println("query userID:", userID)

	var body = UserBody{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		(&httputil.RESP{
			Content:     "",
			ServiceCode: config.Config.Server.Code,
			Err:         fmt.Errorf("body parse failed. err:%v", err),
			HttpCode:    http.StatusBadRequest,
		}).Output(ctx)
		return
	}
	fmt.Println("body:", body)

	ctl := controller.FactoryDemo()

	user, err := ctl.FetchFirst(ctx)
	(&httputil.RESP{
		Content: resp.UserRESP{
			ID:       user.ID.ID,
			Username: user.Username,
			Nickname: user.Nickname,
		},
		ServiceCode: config.Config.Server.Code,
		Err:         err,
	}).Output(ctx)

	return
}
`, moduleName, moduleName))
}
