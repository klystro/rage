package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/klystro/rage/internal/utils"
)

type TemplateData struct {
	ProjectName string
}

func CreateProject(projectName string) error {
	dirs := []string{
		"cmd", "config", "internal/server", "pkg", "modules",
	}

	for _, dir := range dirs {
		path := filepath.Join(projectName, dir)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}

	templateDir, err := utils.GetTemplateDir()
	if err != nil {
		return fmt.Errorf("failed to get template directory: %w", err)
	}

	templates := map[string]string{
		"go.mod":  filepath.Join(templateDir, "base/project.gotmpl"),
		"main.go": filepath.Join(templateDir, "base/main.gotmpl"),
	}

	data := TemplateData{ProjectName: projectName}
	for name, templatePath := range templates {
		path := filepath.Join(projectName, name)
		content, err := utils.RenderTemplate(templatePath, data)
		if err != nil {
			return err
		}
		if err := utils.WriteFile(path, content); err != nil {
			return err
		}
	}

	// Git init & go mod tidy
	_ = os.Chdir(projectName)
	_ = exec.Command("git", "init").Run()
	_ = exec.Command("go", "mod", "tidy").Run()
	return nil
}
