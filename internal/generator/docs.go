package generator

import (
	"fmt"
	"path/filepath"

	"github.com/klystro/rage/internal/utils"
)

func GenerateDocs() error {
	templateDir, err := utils.GetTemplateDir()
	if err != nil {
		return fmt.Errorf("failed to get template directory: %w", err)
	}

	dir := "docs"
	filePath := filepath.Join(dir, "swagger.json")
	templatePath := filepath.Join(templateDir, "docs/swagger.gotmpl")

	content, err := utils.RenderTemplate(templatePath, nil)
	if err != nil {
		return err
	}

	return utils.WriteFile(filePath, content)
}