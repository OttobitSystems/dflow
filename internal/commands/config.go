// Package commands: Terminal UI commands
package commands

import (
	"dflow/internal/cloud/auth"
	"dflow/internal/persistency/repository"
	"fmt"

	"github.com/spf13/cobra"
)

var Config = &cobra.Command{
	Use:   "config",
	Short: "Manage your DevFlow configuration",
}

var SetDefaultFlow = &cobra.Command{
	Use:   "default-flow <flow-name>",
	Short: "Sets default application flow",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repository.UpdateDefaultFlowName(args[0])
	},
}

var SetUserName = &cobra.Command{
	Use:   "username <name>",
	Short: "Sets application username",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repository.UpdateUserName(args[0])
	},
}

var SetCustomerCode = &cobra.Command{
	Use:   "customer <name>",
	Short: "Sets application's customer code",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repository.UpdateClientID(args[0])
	},
}

var Get = &cobra.Command{
	Use:   "get",
	Short: "Gets application's configurations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Username:\t\t %s\n", repository.ApplicationConfiguration.Username)
		fmt.Printf("Default flow name:\t %s\n", repository.ApplicationConfiguration.DefaultFlow)
		fmt.Printf("Joined space:\t\t %s\n", repository.ApplicationConfiguration.JoinedSpace)
		fmt.Printf("Customer code:\t\t %s\n", repository.ApplicationConfiguration.ClientID)

		connectedToCloud := "no"
		if auth.UserLogedInCloud {
			connectedToCloud = "yes"
		}

		fmt.Printf("Connected to cloud:\t %s\n", connectedToCloud)
	},
}
