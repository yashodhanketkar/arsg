package ui

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
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Up, k.Down}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Copy},
		{k.Quit, k.Reset, k.StartOver},
		{k.Help},
	}
}

var keys = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
	Reset: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "reset field"),
	),
	StartOver: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "start over"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Copy: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "copy rating"),
	),
}
