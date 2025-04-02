package ui

import (
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
	scoreMode  int
	score      float32
	help       help.Model
	keys       KeyMap
	view       int
}

func initialModel() model {
	m := model{
		inputs: make([]textinput.Model, 6),
	}

	m.scoreMode = 0
	m.keys = keys
	m.help = help.New()
	m.view = 0

	var t textinput.Model
	for i := range m.inputs {

		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 64

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
