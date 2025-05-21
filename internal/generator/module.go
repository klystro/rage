package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

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
