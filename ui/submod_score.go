package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) scoreUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.ratings.SetWidth(128)

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Batch(tea.ExitAltScreen, tea.Quit)

		case "f3":
			m.view = 0
			return m, nil
		}
	}
	var cmd tea.Cmd
	m.ratings, cmd = m.ratings.Update(msg)
	return m, cmd
}

func (m model) scoreView() string {
	var b strings.Builder
	b.WriteString(defaultStyle.Render(m.ratings.View()))
	return b.String()
}
