package file

import "strings"

type ZRoute struct{}

func (route *ZRoute) Name() string {
	return "z_route.go"
}

func (route *ZRoute) Content() string {
	return strings.TrimSpace(`
package route

import (
    "github.com/gin-gonic/gin"
)

// SetupRoutes sets up all the routes for the API
func SetupRoutes(r *gin.Engine) {
    // Example route
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
}
`)
}
