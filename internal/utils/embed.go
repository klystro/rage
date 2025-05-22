package utils

import (
	"bytes"
	"embed"
	"path/filepath"
	"text/template"
)

// go:embed templates/*.tmpl
var templates embed.FS

// GetTemplateDir returns the absolute path to the embedded templates
func GetTemplateDir() (string, error) {
	return "templates", nil
}

// RenderTemplate updates to use embedded templates
func RenderTemplate(templatePath string, data interface{}) ([]byte, error) {
	tmpl, err := templates.ReadFile(templatePath)
	if err != nil {
		return nil, err
	}

	// Parse and execute template from the embedded content
	t, err := template.New(filepath.Base(templatePath)).Parse(string(tmpl))
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
