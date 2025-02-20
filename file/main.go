package file

import (
	"fmt"
	"github.com/lie-flat-planet/gen-project/global"
	"strings"
)

type Main struct {
}

func (*Main) Name() string {
	return "main.go"
}

func (*Main) Content() string {
	return strings.TrimSpace(fmt.Sprintf(`
package main

// @title %s
// @version 1.0
// @description This is a %s
// @termsOfService http://www.swagger.io/support

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

import (
	"github.com/lie-flat-planet/service-init-tool/command"
	"%s/api/router"
	"%s/config"
	"github.com/spf13/cobra"
)

func init() {
	command.AddCommand(
		&cobra.Command{
			Use: "migrate",
			RunE: func(cmd *cobra.Command, args []string) error {
				return config.Config.Mysql.MigrateAll()
			},
		})
}

func main() {
	command.Execute(func(cmd *cobra.Command, args []string) {
		router.Start()
	})
}

`, global.GetProjectName(), global.GetProjectName(), global.GetModuleName(), global.GetModuleName()))
}
