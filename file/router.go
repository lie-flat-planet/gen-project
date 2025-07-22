package file

import (
	"fmt"
	"github.com/lie-flat-planet/gen-project/global"
	"strings"
)

type Route struct{}

func (route *Route) Name() string {
	return "router.go"
}

func (route *Route) Content() string {
	moduleName := global.GetModuleName()
	projectName := global.GetProjectName()
	return strings.TrimSpace(fmt.Sprintf(`
package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"%s/api/router/demo"
	_ "%s/cmd/%s/docs"
	"%s/config"
)

func NewRoot(r *gin.Engine) {
	basePath := r.Group("/%s/api")
	v1 := basePath.Group("/v1")

	demo.Router(v1)

}

func Start() {
	r := &router{}
	r.init().registerHandler().swagger()

	config.Config.Server.GinServe(r.getEngin())
}

type router struct {
	ginEngine *gin.Engine
}

func (r *router) init() *router {
	gin.SetMode(config.Config.Server.RunMode)
	r.ginEngine = gin.Default()

	return r
}

func (r *router) registerHandler() *router {
	NewRoot(r.ginEngine)

	return r
}

func (r *router) swagger() *router {
	// docs.SwaggerInfo.BasePath = "/%s/api/v1"
	r.ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func (r *router) getEngin() *gin.Engine {
	return r.ginEngine
}
`, moduleName, moduleName, projectName, moduleName, projectName, projectName))
}
