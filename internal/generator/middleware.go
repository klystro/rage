package generator

import (
	"path/filepath"

	"github.com/klystro/rage/internal/utils"
)

func GenerateMiddleware(name string) error {
	dir := filepath.Join("internal", "middleware")
	filePath := filepath.Join(dir, name+".go")
	templatePath := filepath.Join("templates/middleware/middleware.gotmpl")

	content, err := utils.RenderTemplate(templatePath, struct{ Name string }{Name: name})
	if err != nil {
		return err
	}

	return utils.WriteFile(filePath, content)
}