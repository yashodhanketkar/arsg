package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m *model) docUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		margins := headerHeight + footerHeight

		m.viewport = viewport.New(msg.Width, msg.Height-margins)
		m.viewport.YPosition = headerHeight

		_, cmd := m.viewport.Update(msg)

		return m, cmd

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Batch(tea.ExitAltScreen, tea.Quit)

		case "j":
			m.viewport.LineDown(1)

		case "d":
			m.viewport.HalfViewUp()

		case "k":
			m.viewport.LineUp(1)

		case "u":
			m.viewport.HalfViewDown()

		case "home":
			m.viewport.GotoTop()

		case "end":
			m.viewport.GotoBottom()

		case "enter", "f1":
			m.view = m.lastview
			return m, nil
		}
	}

	return m, nil
}

func (m *model) docView() string {
	var b strings.Builder

	b.WriteString(m.headerView() + "\n")
	b.WriteString(m.viewport.View())
	b.WriteString("\n" + m.footerView())

	return b.String()
}

func (m model) headerView() string {
	title := titleStyle.Render("ARGS")
	line := focusedStyle.Render(strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title))))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := focusedStyle.Render(strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info))))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}
