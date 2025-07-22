package command

import (
	"fmt"
	"github.com/lie-flat-planet/gen-project/generator"
	"github.com/lie-flat-planet/gen-project/global"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

const (
	VERSION = "v1.0.13"
)

func init() {
	RootCMD.SetHelpTemplate(`Describe:
  {{.Short}} {{.Version}} - {{.Long}}

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
		PostRunE: func(cmd *cobra.Command, args []string) error { // git init
			// 切换目录
			{
				projectPath := fmt.Sprintf("./%s", global.ParseProjectName(args[0]))
				if err := os.Chdir(projectPath); err != nil {
					return err
				}
				cwd, _ := os.Getwd()
				fmt.Printf("cd the path: %s\n", cwd)
			}

			// git init
			command := exec.Command("git", "init")
			_, err := command.Output()
			if err != nil {
				return err
			}
			return nil
		},
	}
)
