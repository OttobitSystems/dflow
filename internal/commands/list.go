package commands

import (
	"dflow/internal/persistency/repository"
	"fmt"

	"github.com/spf13/cobra"
)

var List = &cobra.Command{
	Use:   "list",
	Short: "List all available flows",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		flows := repository.GetFlows()

		outputString := ""
		for _, entry := range flows {
			outputString += fmt.Sprintf("> Created at: '%s' | Name: %s\n", entry.CreatedAt.Local().Format("2006-01-02 15:04:05"), entry.Name)
		}

		fmt.Println(outputString)
	},
}
