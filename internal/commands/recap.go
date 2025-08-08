package commands

import (
	"dflow/internal/recap"
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var Recap = &cobra.Command{
	Use:   "recap",
	Short: "Shows a recap of work flows",
	Run: func(cmd *cobra.Command, args []string) {
		/*
			fmt.Println(recap.Calculate())

			fmt.Println("NAME\t\t\tLAST ENTER\t\t\tTIME IN FLOW\t\t\tFLOWSPACE")
			for _, entry := range recap.Calculate().FlowsRecap {
				fmt.Printf("%s\t\t\t%s\t\t\t%s\t\t\t\n", entry.Name, entry.LastEnter, entry.TimeInFlow)
			}

		*/

		recap := recap.Calculate()

		columns := []table.Column{
			{"Name", 35},
			{"Last enter", 35},
			{"Time in flow", 35},
			{"Flowspace", 35},
		}

		rows := make([]table.Row, len(recap.FlowsRecap))

		i := 0
		for _, entry := range recap.FlowsRecap {
			timeInFlow := fmt.Sprintf("%02d h %02d m %02d s",
				int(int(entry.TimeInFlow.Seconds())/3600),
				int(int(entry.TimeInFlow.Seconds())%3600/60),
				int(entry.TimeInFlow.Seconds())%60)

			rows[i] = table.Row{
				entry.Name,
				entry.LastEnter.Format("2006-01-02 15:04:05"),
				timeInFlow,
			}
			i++
		}

		t := table.New(
			table.WithColumns(columns),
			table.WithRows(rows),
		)

		t.SetStyles(table.Styles{
			Cell:   lipgloss.NewStyle().Padding(0, 1),
			Header: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#7D56F4")).Padding(0, 1),
		})

		fmt.Println(t.View())
	},
}
