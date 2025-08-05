package tui

import tea "github.com/charmbracelet/bubbletea/v2"

func InitProgram(initModel tea.Model) *tea.Program {
	return tea.NewProgram(initModel)
}
