package tabs

import (
	"cronicle/ui/constants"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	CurrSection int
}

func NewModel() Model {
	return Model{
		CurrSection: 0,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	var tabs []string
	for i, section := range constants.Sections {
		if m.CurrSection == i {
			tabs = append(tabs, activeTab.Render(section))
		} else {
			tabs = append(tabs, tab.Render(section))
		}
	}

	return tabRow.Copy().Render(lipgloss.JoinHorizontal(lipgloss.Top, tabs...))
}

func (m *Model) SetCurrSection(id int) {
	m.CurrSection = id
}
