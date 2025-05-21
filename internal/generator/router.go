package generator

import (
	"path/filepath"

	"github.com/klystro/rage/internal/utils"
)

func GenerateRouter(name string) error {
	dir := filepath.Join("internal", "router")
	filePath := filepath.Join(dir, name+".go")
	templatePath := filepath.Join("internal/generator/templates/router/router.gotmpl")

	content, err := utils.RenderTemplate(templatePath, struct{ Name string }{Name: name})
	if err != nil {
		return err
	}

	return utils.WriteFile(filePath, content)
}