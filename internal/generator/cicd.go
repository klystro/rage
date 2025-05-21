package generator

import (
	"fmt"
	"path/filepath"

	"github.com/klystro/rage/internal/utils"
)

func GenerateCICD(provider string) error {
	var (
		outputPath string
		templatePath string
	)

	switch provider {
	case "github":
		outputPath = filepath.Join(".github", "workflows", "ci.yml")
		templatePath = filepath.Join("internal/generator/templates/infrastructure/cicd/github.gotmpl")
	case "gitlab":
		outputPath = ".gitlab-ci.yml"
		templatePath = filepath.Join("internal/generator/templates/infrastructure/cicd/gitlab.gotmpl")
	default:
		return fmt.Errorf("unsupported provider: %s", provider)
	}

	content, err := utils.RenderTemplate(templatePath, nil)
	if err != nil {
		return err
	}

	return utils.WriteFile(outputPath, content)
}