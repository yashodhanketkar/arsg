package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yashodhanketkar/arsg/util"
)

type model struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
	scoreMode  int
	score      float32
	help       help.Model
	keys       util.KeyMap
	view       int
	ratings    list.Model
}

func initialModel() model {
	m := model{
		inputs: make([]textinput.Model, 6),
	}

	m.scoreMode = 0
	m.keys = util.AppKeys
	m.help = help.New()
	m.view = 0

	var t textinput.Model
	for i := range m.inputs {

		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 128

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
		}

		m.inputs[i] = t
	}

	return m
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen, textinput.Blink)
}

func (m model) Update(msg tea.Msg) (md tea.Model, cmd tea.Cmd) {
	switch m.view {
	case 0:
		md, cmd = m.formUpdate(msg)
	case 1:
		md, cmd = m.confirmUpdate(msg)
	case 2:
		md, cmd = m.scoreUpdate(msg)
	case 3:
		md, cmd = m.docUpdate(msg)
	}

	return md, cmd
}

func (m model) View() (view string) {
	switch m.view {
	case 0:
		view = m.formView()
	case 1:
		view = m.confirmView()
	case 2:
		view = m.scoreView()
	case 3:
		view = m.docView()
	default:
		view = "Error"
	}

	return view
}

func TeaUI() {
	if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
