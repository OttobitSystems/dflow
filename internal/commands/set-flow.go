// Package commands: Terminal UI commands
package commands

import (
	"dflow/internal/persistency/repository"

	"github.com/spf13/cobra"
)

var CreateFlowCommand = &cobra.Command{
	Use:   "flow <flow name>",
	Short: "Creates new flow name (names must be unique)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		flowName := args[0]
		repository.CreateFlow(flowName)
	},
}
