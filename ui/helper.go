package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"github.com/yashodhanketkar/arsg/db"
	"github.com/yashodhanketkar/arsg/util"

	tea "github.com/charmbracelet/bubbletea"
)

type item struct {
	id         int
	title      string
	desc       string
	parameters []float32
	score      string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return fmt.Sprintf("%s - %+v", i.score, i.parameters) }
func (i item) FilterValue() string { return i.title }

var (
	// INFO: Text and input styles
	defaultStyle        = lipgloss.NewStyle().Margin(1, 2)
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
	keymapStyle         = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(0, 1)

	contentStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("62")).
			Foreground(lipgloss.Color("15")).
			Margin(1, 0).
			Padding(0, 1)

	resultStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("120")).
			Background(lipgloss.Color("240")).
			Padding(0, 1).
			Bold(true)

	// INFO: Button styles
	focusedButtonCf  = focusedStyle.Render("[ Confirm ]")
	focusedButtonSv  = focusedStyle.Render("[ Save ]")
	focusedButtonEnd = focusedStyle.Render("[ End ]")
	focusedButtonRes = focusedStyle.Render("[ Restart ]")

	// blurredButtonCf  = fmt.Sprintf("[ %s ]", blurredStyle.Render("Confirm"))
	blurredButtonSv  = fmt.Sprintf("[ %s ]", blurredStyle.Render("Save"))
	blurredButtonEnd = fmt.Sprintf("[ %s ]", blurredStyle.Render("End"))
	blurredButtonRes = fmt.Sprintf("[ %s ]", blurredStyle.Render("Restart"))

	// INFO: Markdown styles
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().
			BorderStyle(b).
			Padding(0, 1).
			Foreground(lipgloss.Color("205")).
			BorderForeground(lipgloss.Color("205"))
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.BorderStyle(b).
			Foreground(lipgloss.Color("205")).
			BorderForeground(lipgloss.Color("205"))
	}()

	//  INFO: Score systems and content types
	scoreSystem = map[int]string{
		0: "Decimal",
		1: "Integer",
		2: "FivePoint",
		3: "Percentage",
	}

	avalCType = []string{"anime", "manga", "lightnovel"}
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
	parameters := make([]float32, len(m.inputs))
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
		if score, err := util.Calculator(parameters...); err == nil {
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
			m.copyToClipboard()
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
			parameters: []float32{rating.Art, rating.Support, rating.Plot, rating.Bias},
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

func (m *model) toggleContentType() string {
	curr := m.contentType
	// circularly iterates through the content types
	m.contentType = avalCType[(slices.Index(avalCType, curr)+1)%len(avalCType)]
	m.buildScoreList()

	return m.contentType
}

func (m *model) buildScoreList() {
	// generates an array of score with respect to the current content type
	m.ratings = list.New(resetScoreList(m.contentType),
		list.NewDefaultDelegate(), 128, 0)
	m.ratings.Title = "Ratings for " + m.contentType
}

func setupParameters(args ...string) []string {
	if len(args) == 0 {
		return []string{"Art/Animation", "Character/Cast", "Plot", "Bias"}
	}

	for i, arg := range args {
		val, err := util.CapitalizeFirstLetter(arg)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		args[i] = val
	}

	return args
}
