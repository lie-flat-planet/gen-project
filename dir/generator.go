package dir

import (
	"fmt"
	"github.com/lie-flat-planet/gen-project/file"
	"os"
	"path/filepath"
)

var Dirs []Dir

// Dir represents a directory structure
type Dir struct {
	Name     string       // Directory name
	Children []Dir        // Nested subdirectories
	Files    []file.IFile // Files within the directory
}

// Create generates the directory, its files, and its children
func (d *Dir) Create(basePath string) error {
	// Create current directory
	currentPath := filepath.Join(basePath, d.Name)
	if err := os.MkdirAll(currentPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", currentPath, err)
	}
	fmt.Println("Created directory:", currentPath)

	// Create files in the current directory
	for _, f := range d.Files {
		filePath := filepath.Join(currentPath, f.Name())
		if err := os.WriteFile(filePath, []byte(f.Content()), 0644); err != nil {
			return fmt.Errorf("failed to create file %s: %w", filePath, err)
		}
		fmt.Println("Created file:", filePath)
	}

	// Recursively create child directories
	for _, child := range d.Children {
		if err := child.Create(currentPath); err != nil {
			return err
		}
	}

	return nil
}
