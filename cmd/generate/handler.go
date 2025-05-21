package generate

import (
	"fmt"
	"os"

	"github.com/klystro/rage/internal/generator"
	"github.com/spf13/cobra"
)

var HandlerCmd = &cobra.Command{
	Use:   "generate handler [module] [name]",
	Short: "Generate CRUD handler in a module",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		module, name := args[0], args[1]
		err := generator.GenerateHandler(module, name)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Printf("Handler '%s' created in module '%s'\n", name, module)
	},
}
