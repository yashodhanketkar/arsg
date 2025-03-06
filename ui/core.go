package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yashodhanketkar/arsg/db"
)

type model struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
	scoreMode  int
	score      float32
	help       help.Model
	keys       KeyMap
}

func initialModel() model {
	m := model{
		inputs: make([]textinput.Model, 6),
	}

	m.scoreMode = 0
	m.keys = keys
	m.help = help.New()

	var t textinput.Model
	for i := range m.inputs {

		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Name"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Art/Animation"
			t.CharLimit = 5
		case 2:
			t.Placeholder = "Character/Cast"
			t.CharLimit = 5
		case 3:
			t.Placeholder = "Plot"
			t.CharLimit = 5
		case 4:
			t.Placeholder = "Bias"
			t.CharLimit = 5
		case 5:
			t.Placeholder = "Comments"
			t.CharLimit = 128
		}

		m.inputs[i] = t
	}

	return m
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen, textinput.Blink)
}

func (m model) isNumeric() bool {
	return m.focusIndex != 0 && m.focusIndex != 5
}

func (m model) prepareRating() db.Rating {

	parseFloat := func(value string) float32 {
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return 0
		}
		pf := float32(f)

		// clamp values
		if pf < 0.0 {
			pf = 0.0
		} else if pf > 10.0 {
			pf = 10.0
		}

		return pf
	}

	return db.Rating{
		Name:     m.inputs[0].Value(),
		Art:      parseFloat(m.inputs[1].Value()),
		Support:  parseFloat(m.inputs[2].Value()),
		Plot:     parseFloat(m.inputs[3].Value()),
		Bias:     parseFloat(m.inputs[4].Value()),
		Comments: m.inputs[5].Value(),
	}
}

func (m model) buttonCommands() (tea.Model, tea.Cmd) {
	l := len(m.inputs)
	switch m.focusIndex {
	case l:
		DB := db.ConnectDB()
		defer DB.Close()
		db.AddRatings(DB, m.prepareRating())
		m.resetInputs()
		return m, nil

	case l + 1:
		m.resetInputs()
		return m, nil

	case l + 2:
		return m, tea.Quit
	}
	return m, nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m model) View() string {
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
		helpStyle.Render("Some shortcut keys won't work in name and comments fields. (c, q, r)"),
	)

	return b.String()
}
