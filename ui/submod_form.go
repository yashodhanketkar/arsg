package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yashodhanketkar/arsg/db"
)

func (m model) formUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c":
			return m, tea.Batch(tea.ExitAltScreen, tea.Quit)

		case "ctrl+e":
			// FIX: Directly calling follwing functions was causing nil pointer error
			// so they are wrapped for now. will fix this issue later
			func() {
				DB := db.ConnectDB()
				defer DB.Close()
				db.ExportData(DB)
			}()
			return m, nil

		case "home":
			m.focusIndex = 0
			return m.setFocus(m.focusIndex)

		case "end":
			m.focusIndex = len(m.inputs) + 2
			return m.setFocus(m.focusIndex)

		case "ctrl+q":
			return m, tea.Batch(tea.ExitAltScreen, tea.Quit)

		case "f3":
			m.ratings = list.New(resetScoreList(), list.NewDefaultDelegate(), 128, 0)
			m.ratings.Title = "Media Name"
			m.view = 2
			m.focusIndex = 0
			m.setFocus(m.focusIndex)
			return m, nil

		// copy score to clipboard
		case "c":
			if !m.isNumeric() {
				break
			}
			m.copyToClipboard()
			return m, nil

		case "ctrl+v":
			m.pasteFromClipbaord()
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
			if m.focusIndex > len(m.inputs)+3 {
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

func (m model) formView() string {
	var b strings.Builder
	helpView := m.help.View(m.keys)

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	saveButton := &blurredButtonSv
	restartButton := &blurredButtonRes
	endButton := &blurredButtonEnd

	if m.focusIndex == len(m.inputs) {
		saveButton = &focusedButtonSv
	}

	if m.focusIndex == len(m.inputs)+1 {
		restartButton = &focusedButtonRes
	}

	if m.focusIndex == len(m.inputs)+2 {
		endButton = &focusedButtonEnd
	}

	b.WriteString("\n\nRating: ")
	b.WriteString(resultStyle.Render(fmt.Sprintf("%.1f", m.score)))

	fmt.Fprintf(&b, "\n\n%s\n", *saveButton)
	fmt.Fprintf(&b, "%s\n", *restartButton)
	fmt.Fprintf(&b, "%s\n\n", *endButton)

	b.WriteString(helpStyle.Render("cursor mode is "))
	b.WriteString(cursorModeHelpStyle.Render(m.cursorMode.String()))
	b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

	b.WriteString(helpStyle.Render("\nscore mode is "))
	b.WriteString(cursorModeHelpStyle.Render(scoreSystem[m.scoreMode]))
	b.WriteString(helpStyle.Render(" system (ctrl+s to change score system)"))

	b.WriteString(helpStyle.Render("\n"))
	b.WriteString(keymapStyle.Render(helpView))
	b.WriteString(helpStyle.Render("\n"))
	b.WriteString(
		helpStyle.Render("Some shortcut keys won't work in name and comments fields. (c, r)"),
	)

	return defaultStyle.Render(b.String())
}
