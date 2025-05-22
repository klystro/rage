package generator

import (
	"github.com/klystro/rage/internal/utils"
)

func GenerateDocker() error {
	templates := map[string]string{
		"Dockerfile":         "templates/infrastructure/docker/Dockerfile.gotmpl",
		".dockerignore":      "templates/infrastructure/docker/dockerignore.gotmpl",
		"docker-compose.yml": "templates/infrastructure/docker/docker-compose.gotmpl",
	}

	for filename, templatePath := range templates {
		content, err := utils.RenderTemplate(templatePath, nil)
		if err != nil {
			return err
		}
		if err := utils.WriteFile(filename, content); err != nil {
			return err
		}
	}
	return nil
}
