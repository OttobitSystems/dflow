package commands

import (
	"dflow/internal/persistency/repository"
	"fmt"

	"github.com/spf13/cobra"
)

var Logs = &cobra.Command{
	Use:   "logs <flow-name*>",
	Short: "List all logs available for flow (flow-name, if no args application uses default one)",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logs := repository.GetLogs("default")

		outputString := "\nFlow Name: default\n"
		sessionID := ""

		for _, entry := range logs {

			if sessionID != entry.Session.ID {
				outputString += fmt.Sprintf("\nSession ID:\t%s\nStarted at:\t%s\nEnded at:\t%s\n\n",
					entry.Session.ID, entry.Session.StartedAt.Local().Format("2006-01-02 15:04:05"), entry.Session.CompletedAt.Local().Format("2006-01-02 15:04:05"))
			}

			sessionID = entry.Session.ID

			outputString += fmt.Sprintf("%s > %s\n",
				entry.TimeStamp.Local().Format("15:04:05"), entry.Log)
		}

		fmt.Println(outputString)
	},
}
