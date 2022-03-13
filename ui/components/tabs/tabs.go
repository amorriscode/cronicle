package tabs

import (
	"cronicle/ui/constants"
	"cronicle/ui/context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	CurrSection int
}

func New() Model {
	return Model{
		CurrSection: 0,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View(ctx context.Context) string {
	var tabs []string

	for i, section := range constants.Sections {
		if m.CurrSection == i {
			tabs = append(tabs, activeTabStyle.Render(section))
		} else {
			tabs = append(tabs, tabStyle.Render(section))
		}
	}

	return tabRowStyle.Copy().Width(ctx.ScreenWidth).Render(lipgloss.JoinHorizontal(lipgloss.Top, tabs...))
}

func (m *Model) SetCurrSection(id int) {
	m.CurrSection = id
}
