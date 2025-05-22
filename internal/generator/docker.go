package generator

import (
	"fmt"
	"path/filepath"

	"github.com/klystro/rage/internal/utils"
)

func GenerateDocker() error {
	templateDir, err := utils.GetTemplateDir()
	if err != nil {
		return fmt.Errorf("failed to get template directory: %w", err)
	}

	templates := map[string]string{
		"Dockerfile":         filepath.Join(templateDir, "infrastructure/docker/Dockerfile.gotmpl"),
		".dockerignore":      filepath.Join(templateDir, "infrastructure/docker/dockerignore.gotmpl"),
		"docker-compose.yml": filepath.Join(templateDir, "infrastructure/docker/docker-compose.gotmpl"),
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
