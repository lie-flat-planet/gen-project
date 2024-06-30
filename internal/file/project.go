package file

import (
	"github.com/lie-flat-planet/gen-project/internal/dir"
	"os"
)

func CreateProjectDir(projectName string) error {
	dirs := dir.GetDirs()

	for _, d := range dirs {
		d.SetProjectName(projectName)

		path := projectName
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}

		if d.SubPath() != "" {
			path = path + "/" + d.SubPath()
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				return err
			}
		}

		if d.Path() != "" {
			path = path + "/" + d.Path()
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				return err
			}
		}

		if d.File() != "" {
			path = path + "/" + d.File()

			if err := os.WriteFile(path, []byte(d.FileContent()), 0644); err != nil {
				return err
			}
		}
	}

	return nil
}
