package commands

import "github.com/spf13/cobra"

var Space = &cobra.Command{
	Use:   "space",
	Short: "Manage your flow space",
	Run:   func(cmd *cobra.Command, args []string) {},
}
