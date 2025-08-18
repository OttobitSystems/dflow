package commands

import (
	auth "dflow/internal/cloud/auth"

	"github.com/spf13/cobra"
)

var Auth = &cobra.Command{
	Use:   "auth",
	Short: "Connects cli to cloud",
	Run: func(cmd *cobra.Command, args []string) {
		auth.LoginWeb()
	},
}
