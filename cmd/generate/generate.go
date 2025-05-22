package generate

import (
	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate various components",
	Long:  `Generate modules, handlers, middleware, and infrastructure components`,
}

func init() {
	GenerateCmd.AddCommand(ModuleCmd)
	GenerateCmd.AddCommand(HandlerCmd)
	GenerateCmd.AddCommand(MiddlewareCmd)
	GenerateCmd.AddCommand(RouterCmd)
	GenerateCmd.AddCommand(DocsCmd)
	GenerateCmd.AddCommand(DockerCmd)
	GenerateCmd.AddCommand(CICDCmd)
}