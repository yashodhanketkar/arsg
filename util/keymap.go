package util

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Up        key.Binding
	Down      key.Binding
	Quit      key.Binding
	Reset     key.Binding
	StartOver key.Binding
	Help      key.Binding
	Copy      key.Binding
	Home      key.Binding
	End       key.Binding
	Export    key.Binding
	Content   key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Content}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Up, k.Down},
		{k.Quit, k.Copy, k.Reset},
		{k.StartOver, k.Home, k.End},
		{k.Export, k.Content},
	}
}

var AppKeys = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("pgup", "up"),
		key.WithHelp("↑/PU", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("pgdown", "Down"),
		key.WithHelp("↓/PD", "move down"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+q"),
		key.WithHelp("ctrl+q", "quit"),
	),
	Reset: key.NewBinding(
		key.WithKeys("del"),
		key.WithHelp("del", "reset field"),
	),
	StartOver: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "start over"),
	),
	Help: key.NewBinding(
		key.WithKeys("F1"),
		key.WithHelp("F1", "toggle help"),
	),
	Copy: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "copy rating"),
	),
	Home: key.NewBinding(
		key.WithKeys("home"),
		key.WithHelp("home", "go to start"),
	),
	End: key.NewBinding(
		key.WithKeys("end"),
		key.WithHelp("end", "go to end"),
	),
	Export: key.NewBinding(
		key.WithKeys("ctrl+e"),
		key.WithHelp("ctrl+e", "export ratings in json format"),
	),
	Content: key.NewBinding(
		key.WithKeys("ctrl+t"),
		key.WithHelp("ctrl+t", "switch content type"),
	),
}
