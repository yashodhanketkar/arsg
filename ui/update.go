package ui

import (
	"github.com/charmbracelet/bubbles/cursor"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) formUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Batch(tea.ExitAltScreen, tea.Quit)

		case "home":
			m.focusIndex = 0
			return m.setFocus(m.focusIndex)

		case "end":
			m.focusIndex = len(m.inputs) + 2
			return m.setFocus(m.focusIndex)

		case "q":
			if !m.isNumeric() {
				break
			}
			return m, tea.Batch(tea.ExitAltScreen, tea.Quit)

		// copy score to clipboard
		case "c":
			if !m.isNumeric() {
				break
			}
			m.scoreToClipboard()
			return m, nil

		case "f1":
			m.help.ShowAll = !m.help.ShowAll

		// reset focused input
		case "delete":
			if m.focusIndex > len(m.inputs)-1 {
				break
			}
			m.inputs[m.focusIndex].Reset()
			return m, nil

		// reset all inputs
		case "esc":
			m.resetInputs()
			return m, nil

		// Change cursor mode
		case "ctrl+r":
			m.cursorMode++
			if m.cursorMode > cursor.CursorHide {
				m.cursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, len(m.inputs))
			for i := range m.inputs {
				cmds[i] = m.inputs[i].Cursor.SetMode(m.cursorMode)
			}
			return m, tea.Batch(cmds...)

		case "ctrl+s":
			m.scoreMode++
			if m.scoreMode > 3 {
				m.scoreMode = 0
			}

		// Set focus to next input
		case " ", "tab", "shift+tab", "enter", "down", "pgdown", "up", "pgup":
			s := msg.String()

			if s == " " {
				if m.focusIndex != 0 && m.focusIndex != 5 {
					return m.buttonCommands()
				} else {
					break
				}
			}

			if s == "enter" {
				return m.buttonCommands()
			}

			// Cycle indexes
			if s == "tab" || s == "down" || s == "pgdown" {
				m.focusIndex++
			}

			// Cycle indexes reverse
			if s == "shift+tab" || s == "up" || s == "pgup" {
				m.focusIndex--
			}

			// handle focus and styles
			if m.focusIndex > len(m.inputs)+2 {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs) + 2
			}

			return m.setFocus(m.focusIndex)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)
	m.calculateScore()

	return m, cmd
}

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

func (m model) Update(msg tea.Msg) (md tea.Model, cmd tea.Cmd) {
	switch m.view {
	case 0:
		md, cmd = m.formUpdate(msg)
	case 1:
		md, cmd = m.confirmUpdate(msg)
	}
	return md, cmd
}
