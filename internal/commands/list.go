package commands

import "github.com/spf13/cobra"

var List = &cobra.Command{
	Use:   "list",
	Short: "List all available flows",
	Run:   func(cmd *cobra.Command, args []string) {},
}
