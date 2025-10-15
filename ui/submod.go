package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yashodhanketkar/arsg/db"
	"github.com/yashodhanketkar/arsg/util"
)

// INFO: Confirm submod

func (m *model) confirmUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Batch(tea.ExitAltScreen, tea.Quit)

		case "f1":
			m.lastview = m.view
			m.view = 3
			return m, nil

		case "enter", " ":
			m.view = 0
			return m, nil
		}
	}

	return m, nil
}

func (m *model) confirmView() string {
	var b strings.Builder

	confirmButton := &focusedButtonCf
	b.WriteString("Ratings saved successfully!\n")
	fmt.Fprintf(&b, "\n%s\n", *confirmButton)

	return defaultStyle.Render(b.String())
}

// INFO: Doc submod

func (m *model) docUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		margins := headerHeight + footerHeight

		m.viewport = viewport.New(msg.Width, msg.Height-margins)
		m.viewport.YPosition = headerHeight

		_, cmd := m.viewport.Update(msg)

		return m, cmd

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Batch(tea.ExitAltScreen, tea.Quit)

		case "j":
			m.viewport.LineDown(1)

		case "d":
			m.viewport.HalfViewUp()

		case "k":
			m.viewport.LineUp(1)

		case "u":
			m.viewport.HalfViewDown()

		case "home":
			m.viewport.GotoTop()

		case "end":
			m.viewport.GotoBottom()

		case "enter", "f1":
			m.view = m.lastview
			return m, nil
		}
	}

	return m, nil
}

func (m *model) docView() string {
	var b strings.Builder

	b.WriteString(m.headerView() + "\n")
	b.WriteString(m.viewport.View())
	b.WriteString("\n" + m.footerView())

	return b.String()
}

func (m model) headerView() string {
	title := titleStyle.Render("ARGS")
	line := focusedStyle.Render(strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title))))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := focusedStyle.Render(strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info))))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

// INFO: Form submod

func (m *model) formUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	config := util.ConfigType{}
	util.LoadConfig(&config)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c":
			return m, tea.Batch(tea.ExitAltScreen, tea.Quit)

		case "ctrl+e":
			// FIX: Directly calling follwing functions was causing nil pointer error
			// so they are wrapped for now. will fix this issue later
			func() {
				DB := db.ConnectDB()
				defer DB.Close()
				db.ExportData(DB, config.ExportPath)
			}()
			return m, nil

		case "home":
			m.focusIndex = 0
			return m.setFocus(m.focusIndex)

		case "end":
			m.focusIndex = len(m.inputs) + 2
			return m.setFocus(m.focusIndex)

		case "ctrl+q":
			return m, tea.Batch(tea.ExitAltScreen, tea.Quit)

		case "ctrl+t":
			m.toggleContentType()
			return m, nil

		case "f3":
			// builds score list for the current content type
			m.buildScoreList()
			m.view = 2
			m.focusIndex = 0
			m.setFocus(m.focusIndex)
			return m, nil

		// copy score to clipboard
		case "c":
			if !m.isNumeric() {
				break
			}
			m.copyToClipboard()
			return m, nil

		case "ctrl+v":
			m.pasteFromClipbaord()
			return m, nil

		case "?":
			m.help.ShowAll = !m.help.ShowAll

		case "f1":
			m.lastview = m.view
			m.view = 3
			return m, nil

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
			if m.focusIndex > len(m.inputs)+3 {
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

func (m *model) formView() string {
	var b strings.Builder
	helpView := m.help.View(m.keys)

	b.WriteString(contentStyle.Render("Running for " + m.contentType))
	b.WriteString(noStyle.Render("\n"))

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
		helpStyle.Render("Some shortcut keys won't work in name and comments fields. (c, r)"),
	)

	return defaultStyle.Render(b.String())
}

// INFO: Score submod

func (m *model) scoreUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.ratings.SetWidth(128)

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Batch(tea.ExitAltScreen, tea.Quit)

		case "f1":
			m.lastview = m.view
			m.view = 3
			return m, nil

		case "ctrl+t":
			m.toggleContentType()
			m.view = 2
			return m, nil

		case "f3":
			m.view = 0
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.ratings, cmd = m.ratings.Update(msg)

	return m, cmd
}

func (m *model) scoreView() string {
	var b strings.Builder
	b.WriteString(defaultStyle.Render(m.ratings.View()))

	return b.String()
}
