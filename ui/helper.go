package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
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

	focusedButton = focusedStyle.Render("[ End ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("End"))

	focusedButtonAlt = focusedStyle.Render("[ Restart ]")
	blurredButtonAlt = fmt.Sprintf("[ %s ]", blurredStyle.Render("Restart"))
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
		m.inputs[i].SetValue(numericInput(m.inputs[i].Value()))
	}

	return tea.Batch(cmds...)

}

func (m *model) calculateScore() {
	var parameters [4]float32
	allValid := true

	for i := range m.inputs {
		if val, err := strconv.ParseFloat(m.inputs[i].Value(), 32); err == nil {
			parameters[i] = float32(val)
		} else {
			allValid = false
		}
	}

	if allValid {
		if score, err := util.Calculator(parameters); err == nil {
			m.score = score
		} else {
			m.score = 0
		}
	} else {
		m.score = 0
	}
}
