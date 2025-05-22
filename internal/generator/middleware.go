package generator

import (
	"fmt"
	"path/filepath"

	"github.com/klystro/rage/internal/utils"
)

func GenerateMiddleware(name string) error {
	templateDir, err := utils.GetTemplateDir()
	if err != nil {
		return fmt.Errorf("failed to get template directory: %w", err)
	}

	dir := filepath.Join("internal", "middleware")
	filePath := filepath.Join(dir, name+".go")
	templatePath := filepath.Join(templateDir, "middleware/middleware.gotmpl")

	content, err := utils.RenderTemplate(templatePath, struct{ Name string }{Name: name})
	if err != nil {
		return err
	}

	return utils.WriteFile(filePath, content)
}
