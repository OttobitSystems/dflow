// Package tui: terminal user interface
package tui

import (
	"dflow/internal/flow"
	"dflow/internal/persistency/repository"
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	footerStyle           = lipgloss.NewStyle().Background(lipgloss.Color("#000000")).Width(100)
	footerLeftContainer   = lipgloss.NewStyle().Background(lipgloss.Color("#7D56F4")).Width(33).Align(lipgloss.Left).Foreground(lipgloss.Color("#FAFAFA"))
	footerCenterContainer = lipgloss.NewStyle().Background(lipgloss.Color("#3C3C3C")).Width(0).Align(lipgloss.Center)
	footerRightContainer  = lipgloss.NewStyle().Background(lipgloss.Color("#D75FEE")).Width(67).Align(lipgloss.Right).Foreground(lipgloss.Color("#FAFAFA"))
)

var (
	leftPanelStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, true, false, false). // Bordo solo a destra
			BorderForeground(lipgloss.Color("63")).
			Padding(1, 2).
			Width(60)

	rightPanelStyle = lipgloss.NewStyle().
			Padding(1, 2)
)

type FlowTickCmd time.Time

type EnterModel struct {
	FlowSession *flow.Session
	InputLog    textinput.Model
}

func (model EnterModel) Init() tea.Cmd {
	return tea.ClearScreen
}

func (model EnterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			if model.FlowSession.IsActive() {
				model.FlowSession.End()
			}
			return model, tea.Quit
		case tea.KeyEnter:
			model.FlowSession.StoreLog(model.InputLog.Value())

			model.InputLog.Reset()

			return model, nil
		default:
			var command tea.Cmd
			model.InputLog, command = model.InputLog.Update(msg)

			return model, command
		}
	}

	// Update logic can go here
	return model, feelItFlow()
}

func (model EnterModel) View() string {
	var view string

	view += fmt.Sprintf("Hi %s, welcome into your flow!\n", repository.ApplicationConfiguration.Username)

	view += "Press `Ctrl+c` to exit\n"

	view += lipgloss.JoinHorizontal(lipgloss.Top, RenderLeftContainer(model), RenderRightContainer(model))

	// ====== FOOTER ========
	view += RenderFooter(model)

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

func RenderLeftContainer(model EnterModel) string {
	var view string

	view = "Logs:\n"
	logs := model.FlowSession.GetLogs()

	for _, entry := range logs {
		view += fmt.Sprintf("> %s\n", entry.Log)
	}

	return leftPanelStyle.Render(view)
}

func RenderRightContainer(model EnterModel) string {
	return rightPanelStyle.Render(model.InputLog.View())
}

func RenderFooter(model EnterModel) string {
	var view string

	leftInfo := fmt.Sprintf("%s > Flow: %s", time.Now().Format("15:04:05"), model.FlowSession.FlowName)
	centerInfo := "" // useless
	rightInfo := fmt.Sprintf("Duration: %s | Objective: %s", model.FlowSession.DurationString(), model.FlowSession.Objective)

	spacing := "\n\n\n\n\n"
	footerRow := lipgloss.JoinHorizontal(lipgloss.Left, footerLeftContainer.Render(leftInfo), footerCenterContainer.Render(centerInfo), footerRightContainer.Render(rightInfo))

	view += spacing
	view += footerStyle.Render(footerRow)

	return view
}
