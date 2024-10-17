package _interface

import (
	"fmt"
	"os"
	"path/filepath"
)

func Create(basePath string, dir IDir) error {
	// Create current directory
	currentPath := filepath.Join(basePath, dir.GetName())
	if err := os.MkdirAll(currentPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", currentPath, err)
	}
	fmt.Println("Created directory:", currentPath)

	for _, f := range dir.Files() {
		filePath := filepath.Join(currentPath, f.Name())
		if err := os.WriteFile(filePath, []byte(f.Content()), 0644); err != nil {
			return fmt.Errorf("failed to create file %s: %w", filePath, err)
		}
		fmt.Println("Created file:", filePath)
	}

	for _, child := range dir.Children() {
		if err := child.Create(currentPath); err != nil {
			return err
		}
	}

	return nil
}
