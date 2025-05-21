package generator

import (
	"fmt"
	"os"
)

func Doctor() error {
	required := []string{"cmd", "modules", "internal/server"}
	for _, r := range required {
		if _, err := os.Stat(r); os.IsNotExist(err) {
			return fmt.Errorf("missing required directory: %s", r)
		}
	}
	return nil
}
