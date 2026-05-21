package ui

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yashodhanketkar/arsg/src/util"
)

var userParameters []string

type FormModel struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
}

type ViewerModel struct {
	docs        string
	viewport    viewport.Model
	contentType string
}

type StateModel struct {
	// Global & router
	view     int
	lastview int
	DB       *sql.DB
	exitMsg  string

	// score
	scoreMode int
	score     float32
	limiter   float32
}

type model struct {
	// state
	state StateModel

	// shared elements
	help help.Model
	keys util.KeyMap

	// internal model
	form    FormModel
	viewer  ViewerModel
	ratings list.Model
}

func initialModel(DB *sql.DB, limiter float32, args ...string) model {
	parameters, err := setupParameters(args...)

	if err != nil {
		m := model{state: StateModel{exitMsg: err.Error()}}
		tea.Quit()
		return m
	}

	m := model{form: FormModel{inputs: make([]textinput.Model, len(parameters)+2)}}

	m.keys = util.AppKeys
	m.help = help.New()

	m.state = StateModel{
		DB:        DB,
		scoreMode: 0,
		view:      0,
		limiter:   limiter,
		exitMsg:   "",
	}

	content, err := m.loadDocs()
	if err != nil {
		m.state.exitMsg = err.Error()
		tea.Quit()
	}

	viewPort := viewport.New(128, 24)
	viewPort.SetContent(content)

	m.viewer = ViewerModel{
		contentType: "anime",
		docs:        content,
		viewport:    viewPort,
	}

	var t textinput.Model
	for i := range m.form.inputs {

		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 128

		switch i {
		case 0:
			t.Placeholder = "Name"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle

		case len(parameters) + 1:
			t.Placeholder = "Comments"

		default:
			t.Placeholder = (parameters)[i-1]
			t.CharLimit = 5
		}

		m.form.inputs[i] = t
	}

	return m
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen, textinput.Blink)
}

func (m model) Update(msg tea.Msg) (md tea.Model, cmd tea.Cmd) {
	switch m.state.view {
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
	switch m.state.view {
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

func TeaUI(db *sql.DB, args ...string) {
	if _, err := tea.NewProgram(initialModel(db, 9.4, args...)).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
