package generate

import (
	"fmt"
	"os"

	"github.com/klystro/rage/internal/generator"
	"github.com/spf13/cobra"
)

var provider string

var CICDCmd = &cobra.Command{
	Use:   "cicd",
	Short: "Generate CI/CD configuration files",
	Run: func(cmd *cobra.Command, args []string) {
		err := generator.GenerateCICD(provider)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Printf("%s CI/CD configuration generated successfully!\n", provider)
	},
}

func init() {
	CICDCmd.Flags().StringVarP(&provider, "provider", "p", "github", "CI/CD provider (github or gitlab)")
}