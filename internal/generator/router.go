package generator

import (
	"fmt"
	"path/filepath"

	"github.com/klystro/rage/internal/utils"
)

func GenerateRouter(name string) error {
	templateDir, err := utils.GetTemplateDir()
	if err != nil {
		return fmt.Errorf("failed to get template directory: %w", err)
	}

	dir := filepath.Join("internal", "router")
	filePath := filepath.Join(dir, name+".go")
	templatePath := filepath.Join(templateDir, "router/router.gotmpl")

	content, err := utils.RenderTemplate(templatePath, struct{ Name string }{Name: name})
	if err != nil {
		return err
	}

	return utils.WriteFile(filePath, content)
}