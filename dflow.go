package main

import (
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
	commands.ConfigSet.AddCommand(commands.CreateFlowCommand)
	commands.Config.AddCommand(commands.ConfigSet)

	// Add subcommands to the root command
	rootCmd.AddCommand(
		commands.Enter,
		commands.List,
		commands.Space,
		commands.Config,
	)
}

func main() {
	repository.InitDatabase()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
