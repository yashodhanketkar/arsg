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
	"github.com/yashodhanketkar/arsg/src/db"
	"github.com/yashodhanketkar/arsg/src/util"

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

const buttonText = "[ %s ]"

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
	blurredButtonSv  = fmt.Sprintf(buttonText, blurredStyle.Render("Save"))
	blurredButtonEnd = fmt.Sprintf(buttonText, blurredStyle.Render("End"))
	blurredButtonRes = fmt.Sprintf(buttonText, blurredStyle.Render("Restart"))

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
	cmds := make([]tea.Cmd, len(m.form.inputs))
	for i := 0; i <= len(m.form.inputs)-1; i++ {
		if i == index {
			// Set focused state
			cmds[i] = m.form.inputs[i].Focus()
			m.form.inputs[i].PromptStyle = focusedStyle
			m.form.inputs[i].TextStyle = focusedStyle
			continue
		}
		// Remove focused state
		m.form.inputs[i].Blur()
		m.form.inputs[i].PromptStyle = noStyle
		m.form.inputs[i].TextStyle = noStyle
	}

	return m, tea.Batch(cmds...)
}

func (m model) isNumeric() bool {
	return m.form.focusIndex > 0 && m.form.focusIndex < 5
}

func (m *model) resetInputs() tea.Cmd {
	cmds := make([]tea.Cmd, len(m.form.inputs))

	for i := range m.form.inputs {
		m.form.inputs[i].Reset()
		m.form.inputs[i].Blur()
		m.form.inputs[i].PromptStyle = noStyle
		m.form.inputs[i].TextStyle = noStyle
	}

	m.form.focusIndex = 0
	m.setFocus(m.form.focusIndex)
	m.calculateScore()

	return tea.Batch(cmds...)
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.form.inputs))

	for i := range m.form.inputs {
		m.form.inputs[i], cmds[i] = m.form.inputs[i].Update(msg)

		switch i {
		case 0, 5:
			m.form.inputs[i].SetValue(m.form.inputs[i].Value())
		default:
			m.form.inputs[i].SetValue(util.GetNumericInput(m.form.inputs[i].Value()))
		}
	}

	return tea.Batch(cmds...)
}

func (m *model) calculateScore() {
	config := util.ConfigType{}
	util.LoadConfig(&config)
	parameters := make([]float32, len(m.form.inputs)-2)
	allValid := true

	for i := range m.form.inputs {
		if i == 0 || i == len(config.Parameters)+1 {
			continue
		}

		if val, err := strconv.ParseFloat(m.form.inputs[i].Value(), 32); err == nil {
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
		m.validator(&config, parameters...)
	} else {
		m.state.score = 0
	}
}

func (m *model) validator(config *util.ConfigType, parameters ...float32) {
	if score, err := util.Calculator(config, m.state.limiter, parameters...); err == nil {
		m.state.score = util.SystemCalculator(scoreSystem[m.state.scoreMode], score)
	} else {
		m.state.score = 0
	}
}

func (m *model) copyToClipboard() {
	clipboard.WriteAll(fmt.Sprintf("%.1f", m.state.score))
}

func (m *model) pasteFromClipbaord() {
	copiedText, err := clipboard.ReadAll()
	if err != nil {
		return
	}

	m.form.inputs[m.form.focusIndex].SetValue(copiedText)
}

func (m model) prepareRating() db.Rating {
	return db.Rating{
		Name:     m.form.inputs[0].Value(),
		Art:      util.FloatParser(m.form.inputs[1].Value()),
		Support:  util.FloatParser(m.form.inputs[2].Value()),
		Plot:     util.FloatParser(m.form.inputs[3].Value()),
		Bias:     util.FloatParser(m.form.inputs[4].Value()),
		Rating:   fmt.Sprintf("%.1f", m.state.score),
		Comments: m.form.inputs[5].Value(),
	}
}

func (m model) buttonCommands() (tea.Model, tea.Cmd) {
	l := len(m.form.inputs)
	switch m.form.focusIndex {
	case l:
		if m.state.score >= 0.1 && m.form.inputs[0].Value() != "" {
			db.AddRatings(m.state.DB, m.prepareRating(), m.viewer.contentType)
			m.copyToClipboard()
			m.resetInputs()
			m.state.view = 1
		} else {
			m.form.focusIndex = 0
			m.setFocus(m.form.focusIndex)
		}
		return m, nil

	case l + 1:
		m.resetInputs()
		return m, nil

	case l + 2:
		return m, tea.Quit

	default:
		m.form.focusIndex++
		m.setFocus(m.form.focusIndex)
		return m, nil
	}
}

func (m *model) resetScoreList(contentType string) []list.Item {
	ratingList := []list.Item{}

	for _, rating := range db.ListRatings(m.state.DB, contentType) {
		ratingList = append(ratingList, item{
			id:         rating.ID,
			title:      rating.Name,
			score:      rating.Rating,
			parameters: []float32{rating.Art, rating.Support, rating.Plot, rating.Bias},
		})
	}

	return ratingList
}

func (m *model) loadDocs() (string, error) {
	path := filepath.Join(os.Getenv("HOME"), ".local/share/args/lib/docs/manual.md")
	content, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

func (m *model) toggleContentType() string {
	curr := m.viewer.contentType
	// circularly iterates through the content types
	m.viewer.contentType = avalCType[(slices.Index(avalCType, curr)+1)%len(avalCType)]
	m.buildScoreList()

	return m.viewer.contentType
}

func (m *model) buildScoreList() {
	// generates an array of score with respect to the current content type
	m.ratings = list.New(m.resetScoreList(m.viewer.contentType),
		list.NewDefaultDelegate(), 128, 0)
	m.ratings.Title = "Ratings for " + m.viewer.contentType
}

func setupParameters(args ...string) ([]string, error) {
	if len(args) == 0 {
		return []string{"Art/Animation", "Character/Cast", "Plot", "Bias"}, nil
	}

	for i, arg := range args {
		val, err := util.CapitalizeFirstLetter(arg)
		if err != nil {
			return nil, err
		}
		args[i] = val
	}

	return args, nil
}
