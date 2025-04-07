package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yashodhanketkar/arsg/db"
)

var defaultStyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
	scoreMode  int
	score      float32
	help       help.Model
	keys       KeyMap
	view       int
	ratings    list.Model
}

type item struct {
	id         int
	title      string
	desc       string
	parameters [4]float32
	score      string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return fmt.Sprintf("%s - %+v", i.score, i.parameters) }
func (i item) FilterValue() string { return i.title }

func resetScoreList() []list.Item {
	ratingList := []list.Item{}
	DB := db.ConnectDB()
	defer DB.Close()

	for _, rating := range db.ListRatings(DB) {
		ratingList = append(ratingList, item{
			id:         rating.ID,
			title:      rating.Name,
			score:      rating.Rating,
			parameters: [4]float32{rating.Art, rating.Support, rating.Plot, rating.Bias},
		})
	}

	return ratingList
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
