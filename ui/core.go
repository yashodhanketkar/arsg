package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
	scoreMode  int
	score      float32
}

func initialModel() model {
	m := model{
		inputs: make([]textinput.Model, 4),
	}

	m.scoreMode = 0

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Art/Animation"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			t.CharLimit = 5
		case 1:
			t.Placeholder = "Character/Cast"
			t.CharLimit = 5
		case 2:
			t.Placeholder = "Plot"
			t.CharLimit = 5
		case 3:
			t.Placeholder = "Bias"
			t.CharLimit = 5
		}

		m.inputs[i] = t
	}

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		// copy score to clipboard
		case "c":
			m.scoreToClipboard()
			return m, nil

		// reset focused input
		case "r":
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
		case "tab", "shift+tab", "enter", "up", "down", "j", "k":
			s := msg.String()

			if s == "enter" && m.focusIndex == len(m.inputs) {
				m.resetInputs()
				return m, nil
			}

			if s == "enter" && m.focusIndex == len(m.inputs)+1 {
				return m, tea.Quit
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" || s == "k" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs)+1 {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs) + 1
			}

			return m.setFocus(m.focusIndex)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)
	m.calculateScore()

	return m, cmd
}

func (m model) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	restartButton := &blurredButtonAlt
	endButton := &blurredButton

	if m.focusIndex == len(m.inputs) {
		restartButton = &focusedButtonAlt
	}

	if m.focusIndex == len(m.inputs)+1 {
		endButton = &focusedButton
	}

	fmt.Fprintf(&b, "\n\nRating: %.1f\n", m.score)
	fmt.Fprintf(&b, "\n\n%s\n", *restartButton)
	fmt.Fprintf(&b, "%s\n", *endButton)

	b.WriteString(helpStyle.Render("cursor mode is "))
	b.WriteString(cursorModeHelpStyle.Render(m.cursorMode.String()))
	b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

	b.WriteString(helpStyle.Render("\nscore mode is "))
	b.WriteString(cursorModeHelpStyle.Render(scoreSystem[m.scoreMode]))
	b.WriteString(helpStyle.Render(" (ctrl+s to change score system)"))

	return b.String()
}
