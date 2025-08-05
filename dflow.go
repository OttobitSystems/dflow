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
	rootCmd.AddCommand(commands.EnterCmd)
	rootCmd.AddCommand(commands.ListCmd)
	rootCmd.AddCommand(commands.SpaceCmd)
	rootCmd.AddCommand(commands.ConfigCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
