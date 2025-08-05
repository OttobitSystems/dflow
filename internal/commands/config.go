package commands

import "github.com/spf13/cobra"

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage your DevFlow configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}
