package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var withTests bool

var generateModuleCmd = &cobra.Command{
	Use:   "generate module [name]",
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

var generateHandlerCmd = &cobra.Command{
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

func init() {
	generateModuleCmd.Flags().BoolVar(&withTests, "with-tests", false, "Include test files")
	rootCmd.AddCommand(generateModuleCmd)
	rootCmd.AddCommand(generateHandlerCmd)
}
