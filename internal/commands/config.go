// Package commands: Terminal UI commands
package commands

import (
	"dflow/internal/persistency/repository"

	"github.com/spf13/cobra"
)

var Config = &cobra.Command{
	Use:   "config",
	Short: "Manage your DevFlow configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var SetDefaultFlow = &cobra.Command{
	Use:   "default-flow <flow-name>",
	Short: "Sets default application flow",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repository.UpdateDefaultFlowName(args[0])
	},
}
