package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/list"
	"github.com/yashodhanketkar/arsg/db"
	"github.com/yashodhanketkar/arsg/util"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *model) setFocus(index int) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, len(m.inputs))
	for i := 0; i <= len(m.inputs)-1; i++ {
		if i == index {
			// Set focused state
			cmds[i] = m.inputs[i].Focus()
			m.inputs[i].PromptStyle = focusedStyle
			m.inputs[i].TextStyle = focusedStyle
			continue
		}
		// Remove focused state
		m.inputs[i].Blur()
		m.inputs[i].PromptStyle = noStyle
		m.inputs[i].TextStyle = noStyle
	}

	return m, tea.Batch(cmds...)
}

func (m model) isNumeric() bool {
	return m.focusIndex > 0 && m.focusIndex < 5
}

func (m *model) resetInputs() tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i].Reset()
		m.inputs[i].Blur()
		m.inputs[i].PromptStyle = noStyle
		m.inputs[i].TextStyle = noStyle
	}

	m.focusIndex = 0
	m.setFocus(m.focusIndex)
	m.calculateScore()

	return tea.Batch(cmds...)
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)

		switch i {
		case 0, 5:
			m.inputs[i].SetValue(m.inputs[i].Value())
		default:
			m.inputs[i].SetValue(util.GetNumericInput(m.inputs[i].Value()))
		}
	}

	return tea.Batch(cmds...)
}

func (m *model) calculateScore() {
	var parameters [4]float32
	allValid := true

	for i := range m.inputs {
		if i == 0 || i == 5 {
			continue
		}

		if val, err := strconv.ParseFloat(m.inputs[i].Value(), 32); err == nil {
			if val > 10.0 {
				parameters[i-1] = float32(10)
			} else {
				parameters[i-1] = float32(val)
			}
		} else {
			allValid = false
		}
	}

	if allValid {
		if score, err := util.Calculator(parameters); err == nil {
			m.score = util.SystemCalculator(scoreSystem[m.scoreMode], score)
		} else {
			m.score = 0
		}
	} else {
		m.score = 0
	}
}

func (m *model) copyToClipboard() {
	clipboard.WriteAll(fmt.Sprintf("%.1f", m.score))
}

func (m *model) pasteFromClipbaord() {
	copiedText, err := clipboard.ReadAll()
	if err != nil {
		return
	}

	m.inputs[m.focusIndex].SetValue(copiedText)
}

func (m model) prepareRating() db.Rating {
	return db.Rating{
		Name:     m.inputs[0].Value(),
		Art:      util.FloatParser(m.inputs[1].Value()),
		Support:  util.FloatParser(m.inputs[2].Value()),
		Plot:     util.FloatParser(m.inputs[3].Value()),
		Bias:     util.FloatParser(m.inputs[4].Value()),
		Rating:   fmt.Sprintf("%.1f", m.score),
		Comments: m.inputs[5].Value(),
	}
}

func (m model) buttonCommands() (tea.Model, tea.Cmd) {
	l := len(m.inputs)
	switch m.focusIndex {
	case l:
		if m.score >= 0.1 && m.inputs[0].Value() != "" {
			DB := db.ConnectDB()
			defer DB.Close()
			db.AddRatings(DB, m.prepareRating(), m.contentType)
			m.resetInputs()
			m.view = 1
		} else {
			m.focusIndex = 0
			m.setFocus(m.focusIndex)
		}
		return m, nil

	case l + 1:
		m.resetInputs()
		return m, nil

	case l + 2:
		return m, tea.Quit

	default:
		m.focusIndex++
		m.setFocus(m.focusIndex)
		return m, nil
	}
}

func resetScoreList(contentType string) []list.Item {
	ratingList := []list.Item{}
	DB := db.ConnectDB()
	defer DB.Close()

	for _, rating := range db.ListRatings(DB, contentType) {
		ratingList = append(ratingList, item{
			id:         rating.ID,
			title:      rating.Name,
			score:      rating.Rating,
			parameters: [4]float32{rating.Art, rating.Support, rating.Plot, rating.Bias},
		})
	}

	return ratingList
}

func (m *model) loadDocs() string {
	path := filepath.Join(os.Getenv("HOME"), ".local/share/args/lib/docs/manual.md")
	content, err := os.ReadFile(path)

	if err != nil {
		fmt.Println("could not load documentation file:", err)
		os.Exit(1)
	}

	return string(content)
}

func (m *model) toggleContentType() {
	curr := m.contentType
	// circularly iterates through the content types
	m.contentType = avalCType[(slices.Index(avalCType, curr)+1)%len(avalCType)]
	m.buildScoreList()
}

func (m *model) buildScoreList() {
	// generates an array of score with respect to the current content type
	m.ratings = list.New(resetScoreList(m.contentType),
		list.NewDefaultDelegate(), 128, 0)
	m.ratings.Title = "Ratings for " + m.contentType
}
