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
	"dflow/internal/persistency/repository"
	"dflow/internal/tui"
	"errors"
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/spf13/cobra"
)

var Objective string

var Enter = &cobra.Command{
	Use:   "enter",
	Short: "Enter the DevFlow",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("too many arguments, expected at most one flow name")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if Objective != "" {
			fmt.Println(Objective)
		}

		p := tui.InitProgram(tui.EnterModel{
			FlowSession: flow.InitSession(setFlowName(args), Objective),
			InputLog:    CreateInputLog(),
		})

		if _, err := p.Run(); err != nil {
			return errors.New("unable to initialize the TUI")
		}
		return nil
	},
}

func CreateInputLog() textinput.Model {
	input := textinput.New()
	input.Placeholder = "Write your log here, press enter to store it!"
	input.Focus()
	input.CharLimit = 150
	input.Width = 40

	return input
}

func setFlowName(args []string) string {
	if len(args) == 0 {
		return repository.ApplicationConfiguration.DefaultFlow
	}
	return args[0]
}
