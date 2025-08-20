package main

import (
	auth "dflow/internal/cloud/auth"
	"dflow/internal/commands"
	"dflow/internal/persistency/repository"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dflow",
	Short: "Welcome to DevFlow CLI",
}

func init() {
	commands.Enter.Flags().StringVarP(&commands.Objective, "objective", "o", "It's a good day today", "Flow objective")

	commands.Config.AddCommand(commands.SetDefaultFlow)
	commands.Config.AddCommand(commands.SetUserName)
	commands.Space.AddCommand(commands.SpaceJoin)
	commands.Space.AddCommand(commands.CreateSpace)
	commands.Space.AddCommand(commands.SpaceRecap)

	// Add subcommands to the root command
	rootCmd.AddCommand(
		commands.Enter,
		commands.List,
		commands.Space,
		commands.Config,
		commands.Logs,
		commands.Recap,
		commands.Auth,
		commands.CreateFlowCommand,
	)
}

func main() {
	auth.RefreshSession()
	repository.InitDatabase()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
