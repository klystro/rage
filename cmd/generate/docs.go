package generate

import (
	"fmt"
	"os"

	"github.com/klystro/rage/internal/generator"
	"github.com/spf13/cobra"
)

var DocsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate API documentation",
	Run: func(cmd *cobra.Command, args []string) {
		err := generator.GenerateDocs()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("API documentation generated successfully!")
	},
}