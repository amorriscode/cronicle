package utils

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Up          key.Binding
	Down        key.Binding
	NextSection key.Binding
	PrevSection key.Binding
	ScrollUp    key.Binding
	ScrollDown  key.Binding
	Help        key.Binding
	Quit        key.Binding
	Escape      key.Binding
	Todo        key.Binding
	Daily       key.Binding
	Brag        key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.ScrollUp, k.ScrollDown},
		{k.PrevSection, k.NextSection},
		{k.Todo, k.Daily, k.Brag},
		{k.Quit},
	}
}

var (
	Keys = KeyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "move down"),
		),
		PrevSection: key.NewBinding(
			key.WithKeys("left", "h"),
			key.WithHelp("←/h", "previous section"),
		),
		NextSection: key.NewBinding(
			key.WithKeys("right", "l"),
			key.WithHelp("→/l", "next section"),
		),
		ScrollUp: key.NewBinding(
			key.WithKeys("ctrl+k"),
			key.WithHelp("ctrl+k", "scroll document up"),
		),
		ScrollDown: key.NewBinding(
			key.WithKeys("ctrl+j"),
			key.WithHelp("ctrl+j", "scroll document down"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
		Escape: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "return to sections"),
		),
		Todo: key.NewBinding(
			key.WithKeys("t"),
			key.WithHelp("t", "create todo"),
		),
		Daily: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "create daily log"),
		),
		Brag: key.NewBinding(
			key.WithKeys("b"),
			key.WithHelp("b", "create brag log"),
		),
	}
)
