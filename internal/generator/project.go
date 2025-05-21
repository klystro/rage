package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
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

	templates := map[string]string{
		"go.mod":  "templates/base/project.gotmpl",
		"main.go": "templates/base/main.gotmpl",
	}

	data := TemplateData{ProjectName: projectName}
	for name, templatePath := range templates {
		path := filepath.Join(projectName, name)
		tmpl, err := template.ParseFiles(filepath.Join("internal/generator", templatePath))
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", templatePath, err)
		}
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		if err := tmpl.Execute(f, data); err != nil {
			return err
		}
	}

	// Git init & go mod tidy
	_ = os.Chdir(projectName)
	_ = exec.Command("git", "init").Run()
	_ = exec.Command("go", "mod", "tidy").Run()
	return nil
}
