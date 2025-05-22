package generator

import (
	"fmt"
	"path/filepath"

	"github.com/klystro/rage/internal/utils"
)

func GenerateCICD(provider string) error {
	var (
		outputPath   string
		templatePath string
	)

	templateDir := "templates" // Using a hardcoded path as a temporary fix
	err := error(nil)
	if err != nil {
		return fmt.Errorf("failed to get template directory: %w", err)
	}

	switch provider {
	case "github":
		outputPath = filepath.Join(".github", "workflows", "ci.yml")
		templatePath = filepath.Join(templateDir, "infrastructure/cicd/github.gotmpl")
	case "gitlab":
		outputPath = ".gitlab-ci.yml"
		templatePath = filepath.Join(templateDir, "infrastructure/cicd/gitlab.gotmpl")
	default:
		return fmt.Errorf("unsupported provider: %s", provider)
	}

	content, err := utils.ReadFile(templatePath)
	if err != nil {
		return err
	}

	return utils.WriteFile(outputPath, content)
}
