package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// WriteFile writes content to a file, creating parent directories if needed
func WriteFile(path string, content []byte) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	if err := os.WriteFile(path, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", path, err)
	}
	return nil
}

// FileExists checks if a file exists and is not a directory
func FileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// DirExists checks if a directory exists
func DirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// ReadFile reads the entire file into memory with better error handling
func ReadFile(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", path, err)
	}
	return content, nil
}

// GetTemplateDir returns the absolute path to the templates directory
// func GetTemplateDir() (string, error) {
// 	// Get the executable path
// 	ex, err := os.Executable()
// 	if err != nil {
// 		return "", err
// 	}

// 	// Get the directory containing the executable
// 	exPath := filepath.Dir(ex)

// 	// Templates are in the same directory as the executable
// 	return filepath.Join(exPath, "templates"), nil
// }
