package generate

import (
	"fmt"
	"os"

	"github.com/klystro/rage/internal/generator"
	"github.com/spf13/cobra"
)

var DockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Generate Docker configuration files",
	Run: func(cmd *cobra.Command, args []string) {
		err := generator.GenerateDocker()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("Docker configuration files generated successfully!")
	},
}