package cmd

import (
	"fmt"

	"github.com/klystro/rage/internal/generator"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check project structure integrity",
	Run: func(cmd *cobra.Command, args []string) {
		if err := generator.Doctor(); err != nil {
			fmt.Println("Project check failed:", err)
		} else {
			fmt.Println("All good. Project structure is healthy!")
		}
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
