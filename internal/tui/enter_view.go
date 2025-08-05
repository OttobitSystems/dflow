package tui

import (
	"dflow/internal/flow"
	"fmt"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
	"time"
)

var (
	durationStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205"))
	footerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			Italic(true)
)

type FlowTickCmd time.Time

type EnterModel struct {
	FlowSession *flow.Session
	FlowLog     FlowStateLogModel
}

func (model EnterModel) Init() tea.Cmd {
	//return tea.ClearScreen
	//return nil
	return tea.ClearScreen
}

func (model EnterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "s", "start":
			// Start the flow state
			model.StartFlow()
			return model, feelItFlow()
		case "q", "ctrl+c":
			return model, tea.Quit
		default:
			return model.UpdateFlowState(msg)
		}
	case FlowTickCmd:
		return model, feelItFlow()
	}

	// Update logic can go here
	return model, nil
}

func (model EnterModel) View() string {
	var view string

	//view += "Welcome to the Main View!\n"

	view += fmt.Sprintf("You are in the flow state: %s\n", model.FlowSession.FlowName)

	// Place SomethingChanging in the view
	view += fmt.Sprintf(durationStyle.Render("Duration: %d \n"), model.FlowSession.DurationInSeconds())
	view += model.FlowLog.View()
	view += footerStyle.Render("\nPress 'q' or 'ctrl+c' to quit.\n")

	return view
}

func (model EnterModel) StartFlow() {
	model.FlowSession.Start()
}

func (model EnterModel) UpdateFlowState(msg tea.Msg) (tea.Model, tea.Cmd) {
	updated, cm := model.FlowLog.Update(msg)
	model.FlowLog = updated.(FlowStateLogModel)
	return model, cm
}

func feelItFlow() tea.Cmd {
	return tea.Tick(1*time.Second, func(t time.Time) tea.Msg {
		return FlowTickCmd(t)
	})
}
