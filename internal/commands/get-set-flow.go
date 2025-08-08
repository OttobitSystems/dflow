// Package commands: Terminal UI commands
package commands

import (
	"dflow/internal/persistency/repository"
	"fmt"

	"github.com/spf13/cobra"
)

var CreateFlowCommand = &cobra.Command{
	Use:   "flow <flow name>",
	Short: "Creates new flow name (names must be unique)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		flowName := args[0]
		repository.CreateFlow(flowName)
	},
}

var GetFlowsCommand = &cobra.Command{
	Use:   "flows",
	Short: "Gets each flow names in configuration",
	Run: func(cmd *cobra.Command, args []string) {
		flows := repository.GetFlows()

		outputString := ""
		for _, entry := range flows {
			outputString += fmt.Sprintf("> Created at: '%s' | Name: %s\n", entry.CreatedAt.Local().Format("2006-01-02 15:04:05"), entry.Name)
		}

		fmt.Println(outputString)
	},
}
