package file

import (
	"github.com/lie-flat-planet/gen-project/internal/dir"
	"os"
)

func CreateProjectDir(projectName string) error {
	dirs := dir.GetDirs()

	for _, d := range dirs {
		d.SetProjectName(projectName)

		if d.RootPath() != "" {
			if err := os.MkdirAll(d.RootPath(), os.ModePerm); err != nil {
				return err
			}
		}

		if d.Path() != "" {
			if err := os.MkdirAll(d.RootPath()+d.Path(), os.ModePerm); err != nil {
				return err
			}
		}

		if d.File() != "" {
			filePath := d.RootPath() + d.Path() + d.File()

			if err := os.WriteFile(filePath, []byte(d.FileContent()), 0644); err != nil {
				return err
			}
		}
	}

	return nil
}
