package command

import (
	"fmt"
	"github.com/lie-flat-planet/gen-project/generator"
	"github.com/spf13/cobra"
)

const (
	VERSION = "v1.0.0"
)

func init() {
	RootCMD.SetHelpTemplate(`Describe:
  {{.Short}} - {{.Long}}

Usage:
  ./{{.Short}} moduleName

{{if .HasExample}}Examples:
  {{.Example}}{{end}}

Use "-h or --help" for more information about a command.
`)

}

var (
	RootCMD = &cobra.Command{
		Use:     "",
		Short:   "gen-project",
		Long:    "automatically generate a go mvc web project",
		Version: VERSION,
		Example: `./gen-project github.com/your-group/your-project-name`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("缺少项目名参数")
			}

			if args[0] == "" {
				return fmt.Errorf("项目名参数不能为空")
			}

			if err := generator.Generate("./", args[0]); err != nil {
				panic(err)
			}

			return nil
		},
	}
)
