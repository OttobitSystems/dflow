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
	Use:   "join <client_id> <space_name>",
	Short: "Joins cloud space",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		clientID := args[0]
		spaceID := args[1]

		token := auth.RefreshSession()
		joinResponse := repository.JoinCloudSpace(token, spaceID, clientID)
		if joinResponse {
			repository.UpdateClientID(clientID)
			repository.UpdateSpaceID(spaceID)
			return
		}
		fmt.Println("Space id not found in cloud")
	},
}

var CreateSpace = &cobra.Command{
	Use:   "create <client_id> <space_name>",
	Short: "Creates cloud space",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		clientID := args[0]
		spaceID := args[1]

		token := auth.RefreshSession()
		creationResponse := repository.AddCloudSpace(token, spaceID, clientID)
		fmt.Println(creationResponse)
	},
}

var RecapSpace = &cobra.Command{
	Use:   "recap",
	Short: "Makes recap of work from cloud",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
