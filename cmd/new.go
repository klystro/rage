package cmd

import (
	"fmt"
	"os"

	"github.com/klystro/rage/internal/generator"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new Go modular project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		if err := generator.CreateProject(projectName); err != nil {
			cmd.PrintErrln("Error creating project:", err)
			os.Exit(1)
			return
		}
		fmt.Println("Project created successfully!")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	// newCmd.Flags().StringP("module", "m", "", "Module name for the new project")
	// newCmd.Flags().StringP("path", "p", "", "Path to create the new project")
	// newCmd.Flags().BoolP("force", "f", false, "Force overwrite existing files")
}
