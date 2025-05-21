package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func GenerateHandler(module, name string) error {
	dir := filepath.Join("modules", module, "handler")
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	filePath := filepath.Join(dir, fmt.Sprintf("%s.go", name))

	tmpl, err := template.ParseFiles(filepath.Join("internal/generator/templates/module/handler.gotmpl"))
	if err != nil {
		return fmt.Errorf("failed to parse handler template: %w", err)
	}

	data := struct{ Name string }{Name: name}
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, data)
}
