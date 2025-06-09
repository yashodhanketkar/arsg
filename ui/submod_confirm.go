package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) confirmUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Batch(tea.ExitAltScreen, tea.Quit)

		case "enter", " ":
			m.view = 0
			return m, nil
		}
	}
	return m, nil
}

func (m model) confirmView() string {
	var b strings.Builder

	confirmButton := &focusedButtonCf
	b.WriteString("Ratings saved successfully!\n")
	fmt.Fprintf(&b, "\n%s\n", *confirmButton)

	return defaultStyle.Render(b.String())
}
