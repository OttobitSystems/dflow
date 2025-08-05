/*
	The Enter command allows a user to enter it's "development flow" or "DevFlow".

	The idea is that through this command, a user can start a flow session.
    The flow session tracks the user's current development state, such as the current project,
	the current task, and any other relevant information that can help in managing the development process.

	Instead of referring to a project, the user is entering a flow state, which has a name.
*/

package commands

import (
	"dflow/internal/flow"
	"dflow/internal/tui"
	"errors"
	"github.com/spf13/cobra"
)

var flowName = "default-flow"

var EnterCmd = &cobra.Command{
	Use:   "enter",
	Short: "Enter the DevFlow",
	RunE:  ExecuteEnter,
}

func ExecuteEnter(cmd *cobra.Command, args []string) error {
	// Logic to handle entering a flow state
	// This could involve setting up the flow state, initializing resources, etc.
	// For now, we will just print a message

	// Create a new flow state
	flowState := &flow.FlowState{FlowName: flowName}

	// Create a new TUI program
	p := tui.InitProgram(tui.EnterModel{Flow: flowState})
	if _, err := p.Run(); err != nil {
		return errors.New("unable to initialize the TUI")
	}
	return nil
}
