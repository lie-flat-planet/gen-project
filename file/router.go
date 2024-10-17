package file

import "strings"

type Route struct{}

func (route *Route) Name() string {
	return "z_route.go"
}

func (route *Route) Content() string {
	return strings.TrimSpace(`
package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"test-gen-project/api/router/demo"
	"test-gen-project/cmd/test-gen-project/docs"
	"test-gen-project/config"
)

func NewRoot(r *gin.Engine) {
	basePath := r.Group("/demo/api")
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
	docs.SwaggerInfo.BasePath = "/demo/api/v1"
	r.ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func (r *router) getEngin() *gin.Engine {
	return r.ginEngine
}
`)
}
