package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

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

	blurredButtonCf  = fmt.Sprintf("[ %s ]", blurredStyle.Render("Confirm"))
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
