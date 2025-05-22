package generator

import (
	"path/filepath"

	"github.com/klystro/rage/internal/utils"
)

func GenerateDocs() error {
	dir := "docs"
	filePath := filepath.Join(dir, "swagger.json")
	templatePath := filepath.Join("templates/docs/swagger.gotmpl")

	content, err := utils.RenderTemplate(templatePath, nil)
	if err != nil {
		return err
	}

	return utils.WriteFile(filePath, content)
}