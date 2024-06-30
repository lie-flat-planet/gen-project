package cmd

import (
	"github.com/lie-flat-planet/gen-project/internal/file"
	"github.com/spf13/cobra"
)

const (
	VERSION = "v1.0.0"
)

var (
	DirCMD = &cobra.Command{
		Use:     "dir-tools",
		Short:   "生成工程目录结构",
		Version: VERSION,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("缺少项目名参数")
			}

			if args[0] == "" {
				panic("项目名参数不能为空")
			}

			if err := file.CreateProjectDir(args[0]); err != nil {
				panic(err)
			}
		},
	}
)
