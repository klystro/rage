package generate

import (
	"fmt"
	"os"

	"github.com/klystro/rage/internal/generator"
	"github.com/spf13/cobra"
)

var withTests bool

var ModuleCmd = &cobra.Command{
	Use:   "module [name]",
	Short: "Generate a new module with layered structure",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		err := generator.GenerateModule(name, withTests)
		if err != nil {
			cmd.PrintErrln("Error creating project:", err)
			os.Exit(1)
		}
		fmt.Printf("Module '%s' created successfully!\n", name)
	},
}

func init() {
	ModuleCmd.Flags().BoolVar(&withTests, "with-tests", false, "Include test files")
}
