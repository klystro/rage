package generate

import (
	"fmt"
	"os"

	"github.com/klystro/rage/internal/generator"
	"github.com/spf13/cobra"
)

var RouterCmd = &cobra.Command{
	Use:   "router [name]",
	Short: "Generate a new router",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		err := generator.GenerateRouter(name)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Printf("Router '%s' created successfully!\n", name)
	},
}