package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type FlowStateLogModel struct {
	FlowObjective string
	RequestLog    bool
	LogInput      textinput.Model
}

func (m FlowStateLogModel) Init() tea.Cmd {
	return nil
}

func (m FlowStateLogModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.LogInput.Focused() {
		var cmd tea.Cmd
		// Handle text input updates
		m.LogInput, cmd = m.LogInput.Update(msg)
		return m, cmd
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "l", "log":
			// Toggle the request log
			m.RequestLog = !m.RequestLog
			m.LogInput.Focus()
			return m, textinput.Blink
		}
	}
	return m, nil
}

func (m FlowStateLogModel) View() string {
	var view string

	if m.RequestLog {
		view += "\n"
		view += "Flow State Log:\n"
		view += fmt.Sprintf("Your current objective is %s \n", m.FlowObjective)
		view += m.LogInput.View()
		view += "\n"
	}

	return view
}
