package dir

import (
	"github.com/lie-flat-planet/gen-project/file"
)

func Generate(path, projectName string) error {
	projectStructure := getProjectStructure(projectName)

	// Start generating directories and files from the current directory
	if err := projectStructure.Create(path); err != nil {
		return err
	}

	return nil
}

// Define the project structure
func getProjectStructure(name string) *Dir {
	return &Dir{
		Name: name,
		Children: []Dir{
			{
				Name: "api",
				Children: []Dir{
					{Name: "controller"},
					{Name: "request"},
					{Name: "response"},
					{
						Name: "route",
						Files: []file.IFile{
							&file.ZRoute{},
						},
					},
				},
			},
			{
				Name: "cmd",
			},
			{
				Name: "internal",
				Children: []Dir{
					{Name: "model"},
					{Name: "service"},
				},
			},
		},
	}
}

// Example content for the route file
//func getRouteFileContent() string {
//	return strings.TrimSpace(`
//package route
//
//import (
//    "github.com/gin-gonic/gin"
//)
//
//// SetupRoutes sets up all the routes for the API
//func SetupRoutes(r *gin.Engine) {
//    // Example route
//    r.GET("/ping", func(c *gin.Context) {
//        c.JSON(200, gin.H{
//            "message": "pong",
//        })
//    })
//}
//`)
//}
