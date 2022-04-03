package pages

import tea "github.com/charmbracelet/bubbletea"

type ChangePageMsg struct {
	page string
}

func ChangePageCmd(page string) tea.Cmd {
	return func() tea.Msg {
		return ChangePageMsg{
			page: page,
		}
	}
}
