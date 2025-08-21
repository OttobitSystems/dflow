package commands

import (
	"dflow/internal/cloud/auth"
	"dflow/internal/persistency/repository"
	"fmt"

	"github.com/spf13/cobra"
)

var Space = &cobra.Command{
	Use:   "space",
	Short: "Manage your flow space",
}

var SpaceJoin = &cobra.Command{
	Use:   "join <space_name>",
	Short: "Joins cloud space",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		clientID := repository.ApplicationConfiguration.ClientID
		spaceID := args[0]

		token := auth.RefreshSession()
		joinResponse := repository.JoinCloudSpace(token, spaceID, clientID)
		if joinResponse {
			repository.UpdateSpaceID(spaceID)
			fmt.Println("Joined")
			return
		}
		fmt.Println("Space id not found in cloud")
	},
}

var CreateSpace = &cobra.Command{
	Use:   "create <space_name>",
	Short: "Creates cloud space",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		clientID := repository.ApplicationConfiguration.ClientID
		spaceID := args[0]

		token := auth.RefreshSession()
		creationResponse := repository.AddCloudSpace(token, spaceID, clientID)
		fmt.Println(creationResponse)
	},
}

var ListSpaces = &cobra.Command{
	Use:   "list",
	Short: "Lists spaces from cloud",
	Run: func(cmd *cobra.Command, args []string) {
		token := auth.RefreshSession()
		response := repository.ListSpaces(token, repository.ApplicationConfiguration.ClientID)
		for _, item := range *response {
			fmt.Printf("> %s\n", item)
		}
	},
}
