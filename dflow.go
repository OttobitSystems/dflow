package main

import (
	"dflow/internal/commands"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "dflow",
	Short: "Welcome to DevFlow CLI",
}

func init() {
	// Add subcommands to the root command
	rootCmd.AddCommand(
		commands.Enter,
		commands.List,
		commands.Space,
		commands.Config,
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
