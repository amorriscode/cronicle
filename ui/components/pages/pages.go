package pages

import tea "github.com/charmbracelet/bubbletea"

type ChangePageMsg struct {
	Page string
}

func ChangePageCmd(page string) tea.Cmd {
	return func() tea.Msg {
		return ChangePageMsg{
			Page: page,
		}
	}
}
