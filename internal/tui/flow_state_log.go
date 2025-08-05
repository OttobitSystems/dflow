package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type FlowStateLogModel struct {
	FlowObjective string
	RequestLog    bool
	//logInput      textinput.Model
}

func (m FlowStateLogModel) Init() tea.Cmd {
	//return textinput.Blink
	return nil
}

func (m FlowStateLogModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "l", "log":
			// Toggle the request log
			m.RequestLog = !m.RequestLog
			return m, nil
		}
	}
	return m, nil
}

func (m FlowStateLogModel) View() string {
	var view string

	if m.RequestLog {
		view += "\n"
		view += "Flow State Log:\n"
		view += fmt.Sprintf("Your current objective is %s", m.FlowObjective)
		view += "\n"
	}

	return view
}
