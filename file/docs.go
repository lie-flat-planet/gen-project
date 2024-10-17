package file

import "strings"

type Docs struct{}

func (*Docs) Name() string {
	return "docs.go"
}

func (*Docs) Content() string {
	return strings.TrimSpace(`
package docs

import "github.com/swaggo/swag"
const docTemplate = ` + `{
	"schemes": {{ marshal .Schemes }},
	"swagger": "2.0",
		"info": {
		"description": "{{escape .Description}}",
			"title": "{{.Title}}",
			"contact": {},
		"version": "{{.Version}}"
	},
	"host": "{{.Host}}",
		"basePath": "{{.BasePath}}",
		"paths": {}
}` + `
// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
`)
}
