package tui

import (
	"dflow/internal/flow"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
}

func (model EnterModel) Init() tea.Cmd {
	return tea.ClearScreen
}

func (model EnterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return model, tea.Quit
		case "s":
			if model.FlowSession.IsActive() {
				// If the flow is already active, end it
				model.FlowSession.End()
			} else {
				// If the flow is not active, start it
				model.FlowSession.Start()
			}
			return model, feelItFlow()
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

	// If the session completed, show the end time
	if model.FlowSession.IsCompleted() {
		view += fmt.Sprintf("Flow ended at: %s\n", model.FlowSession.EndedAt.Format(time.RFC1123))
	} else {
		view += fmt.Sprintf("Flow started at: %s\n", model.FlowSession.StartedAt.Format(time.RFC1123))
	}

	view += footerStyle.Render("\nPress 'q' or 'ctrl+c' to quit.\n")

	return view
}

func (model EnterModel) StartFlow() {
	model.FlowSession.Start()
}

func feelItFlow() tea.Cmd {
	return tea.Tick(1*time.Second, func(t time.Time) tea.Msg {
		return FlowTickCmd(t)
	})
}
