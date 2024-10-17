package file

import (
	"fmt"
	"strings"

	"github.com/lie-flat-planet/gen-project/global"
)

type ControllerDemo struct{}

func (*ControllerDemo) Name() string {
	return "demo.go"
}

func (*ControllerDemo) Content() string {
	moduleName := global.GetModuleName()

	return strings.TrimSpace(fmt.Sprintf(`
package controller

import (
	"context"
	"%s/config"
	"%s/internal/model"
	"%s/internal/service"
)

func FactoryDemo() *Demo {
	return &Demo{
		service: service.NewDemo(
			service.WithMysql(config.Config.Mysql.GetDB()),
			service.WithRedis(config.Config.Redis.GetClient()),
		),
	}
}

type Demo struct {
	service *service.Demo
}

func (demo *Demo) FetchFirst(ctx context.Context) (*model.User, error) {
	return demo.service.FetchFirst(ctx)
}

`, moduleName, moduleName, moduleName))
}