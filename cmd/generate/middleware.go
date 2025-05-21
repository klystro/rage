package generate

import (
	"fmt"
	"os"

	"github.com/klystro/rage/internal/generator"
	"github.com/spf13/cobra"
)

var MiddlewareCmd = &cobra.Command{
	Use:   "middleware [name]",
	Short: "Generate a new middleware",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		err := generator.GenerateMiddleware(name)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Printf("Middleware '%s' created successfully!\n", name)
	},
}