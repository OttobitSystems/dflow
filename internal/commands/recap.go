package commands

import (
	"dflow/internal/cloud/auth"
	"dflow/internal/persistency/repository"
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
		recap := recap.Calculate()

		columns := []table.Column{
			{"Name", 20},
			{"Last enter", 20},
			{"Time in flow", 20},
			{"Flowspace", 20},
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

var SpaceRecap = &cobra.Command{
	Use:   "recap",
	Short: "Shows a recap of work flows",
	Run: func(cmd *cobra.Command, args []string) {
		token := auth.RefreshSession()
		flow := repository.RetriveRecapData(token,
			repository.ApplicationConfiguration.JoinedSpace,
			repository.ApplicationConfiguration.ClientID,
			repository.ApplicationConfiguration.DefaultFlow)

		recap := recap.CalculateWithFlow(*flow)

		columns := []table.Column{
			{"Name", 20},
			{"Last enter", 20},
			{"Time in flow", 20},
			{"Flowspace", 20},
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
