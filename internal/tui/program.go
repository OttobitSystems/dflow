package tui

import tea "github.com/charmbracelet/bubbletea"

func InitProgram(initModel tea.Model) *tea.Program {
	return tea.NewProgram(initModel)
}
