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

	files := map[string]string{
		"go.mod":  `module {{.ProjectName}}\n\ngo 1.20`,
		"main.go": `package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello from {{.ProjectName}}!\")\n}`,
	}

	data := TemplateData{ProjectName: projectName}
	for name, content := range files {
		path := filepath.Join(projectName, name)
		tmpl, err := template.New(name).Parse(content)
		if err != nil {
			return err
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

func GenerateModule(name string, withTests bool) error {
	modulePath := filepath.Join("modules", name)
	structure := []string{
		"handler", "model", "repository", "usecase", "service",
	}

	for _, layer := range structure {
		dir := filepath.Join(modulePath, layer)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
		filePath := filepath.Join(dir, fmt.Sprintf("%s.go", layer))
		fileContent := fmt.Sprintf("package %s\n\n// TODO: implement %s logic", layer, layer)
		if err := os.WriteFile(filePath, []byte(fileContent), 0644); err != nil {
			return err
		}
		if withTests {
			testFilePath := filepath.Join(dir, fmt.Sprintf("%s_test.go", layer))
			testContent := fmt.Sprintf("package %s\n\nimport \"testing\"\n\nfunc TestDummy(t *testing.T) {\n\tt.Log(\"%s test\")\n}", layer, layer)
			if err := os.WriteFile(testFilePath, []byte(testContent), 0644); err != nil {
				return err
			}
		}
	}

	if name == "auth" {
		authUtil := filepath.Join(modulePath, "util")
		_ = os.MkdirAll(authUtil, os.ModePerm)
		_ = os.WriteFile(filepath.Join(authUtil, "jwt.go"), []byte("package util\n\n// TODO: implement JWT utils"), 0644)
	}

	if name == "cache" {
		cacheUtil := filepath.Join(modulePath, "util")
		_ = os.MkdirAll(cacheUtil, os.ModePerm)
		_ = os.WriteFile(filepath.Join(cacheUtil, "redis.go"), []byte("package util\n\n// TODO: implement Redis caching"), 0644)
	}

	return nil
}

func GenerateHandler(module, name string) error {
	dir := filepath.Join("modules", module, "handler")
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	filePath := filepath.Join(dir, fmt.Sprintf("%s.go", name))
	tmpl := `package handler

import "net/http"

func Create{{.Name}}(w http.ResponseWriter, r *http.Request) {}
func Get{{.Name}}(w http.ResponseWriter, r *http.Request) {}
func Update{{.Name}}(w http.ResponseWriter, r *http.Request) {}
func Delete{{.Name}}(w http.ResponseWriter, r *http.Request) {}
`
	data := struct{ Name string }{Name: name}
	t, _ := template.New("crud").Parse(tmpl)
	f, _ := os.Create(filePath)
	defer f.Close()
	t.Execute(f, data)
	return nil
}

func Doctor() error {
	required := []string{"cmd", "modules", "internal/server"}
	for _, r := range required {
		if _, err := os.Stat(r); os.IsNotExist(err) {
			return fmt.Errorf("missing required directory: %s", r)
		}
	}
	return nil
}
