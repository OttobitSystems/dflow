package commands

import "github.com/spf13/cobra"

var Config = &cobra.Command{
	Use:   "config",
	Short: "Manage your DevFlow configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}
