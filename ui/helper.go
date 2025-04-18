package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/lipgloss"
	"github.com/yashodhanketkar/arsg/db"
	"github.com/yashodhanketkar/arsg/util"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	resultStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("120")).
			Background(lipgloss.Color("240")).Padding(0, 1).Bold(true)

	keymapStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).Padding(0, 1)

	focusedButtonCf = focusedStyle.Render("[ Confirm ]")
	blurredButtonCf = fmt.Sprintf("[ %s ]", blurredStyle.Render("Confirm"))

	focusedButtonSv = focusedStyle.Render("[ Save ]")
	blurredButtonSv = fmt.Sprintf("[ %s ]", blurredStyle.Render("Save"))

	focusedButtonEnd = focusedStyle.Render("[ End ]")
	blurredButtonEnd = fmt.Sprintf("[ %s ]", blurredStyle.Render("End"))

	focusedButtonRes = focusedStyle.Render("[ Restart ]")
	blurredButtonRes = fmt.Sprintf("[ %s ]", blurredStyle.Render("Restart"))

	scoreSystem = map[int]string{
		0: "Decimal",
		1: "Integer",
		2: "FivePoint",
		3: "Percentage",
	}
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
	return m.focusIndex != 0 && m.focusIndex != 5
}

func numericInput(str string) string {
	var inputBuilder strings.Builder

	for _, r := range str {
		if strings.ContainsRune("0123456789.", r) {
			inputBuilder.WriteRune(r)
		}
	}
	return inputBuilder.String()
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
			m.inputs[i].SetValue(numericInput(m.inputs[i].Value()))
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

func (m *model) scoreToClipboard() {
	clipboard.WriteAll(fmt.Sprintf("%.1f", m.score))
}

func (m *model) scoreFromClipbaord() {
	copiedText, err := clipboard.ReadAll()
	if err != nil {
		return
	}

	m.inputs[m.focusIndex].SetValue(copiedText)
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
			db.AddRatings(DB, m.prepareRating())
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
