package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/klystro/rage/internal/utils"
)

func GenerateHandler(module, name string) error {
	templateDir, err := utils.GetTemplateDir()
	if err != nil {
		return fmt.Errorf("failed to get template directory: %w", err)
	}

	dir := filepath.Join("modules", module, "handler")
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	filePath := filepath.Join(dir, fmt.Sprintf("%s.go", name))
	templatePath := filepath.Join(templateDir, "module/handler.gotmpl")

	content, err := utils.RenderTemplate(templatePath, struct{ Name string }{Name: name})
	if err != nil {
		return err
	}

	return utils.WriteFile(filePath, content)
}
