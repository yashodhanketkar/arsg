package ui

import tea "github.com/charmbracelet/bubbletea"

func (m model) docUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) docView() string {
	return ""
}
