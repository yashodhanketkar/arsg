package ui

import (
	"fmt"
	"strings"
)

func (m model) formView() string {
	var b strings.Builder
	helpView := m.help.View(m.keys)

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
		helpStyle.Render("Some shortcut keys won't work in name and comments fields. (c, q, r)"),
	)

	return b.String()
}

func (m model) confirmView() string {
	var b strings.Builder

	confirmButton := &focusedButtonCf
	b.WriteString("Ratings saved successfully!\n")
	fmt.Fprintf(&b, "\n%s\n", *confirmButton)

	return b.String()
}

func (m model) View() (view string) {
	switch m.view {
	case 0:
		view = m.formView()
	case 1:
		view = m.confirmView()
	default:
		view = "Error"
	}

	return view
}
