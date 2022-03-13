package help

import (
	"cronicle/ui/context"
	"cronicle/utils"

	bubblesHelp "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	help bubblesHelp.Model
}

func New() Model {
	help := bubblesHelp.New()

	return Model{
		help: help,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, utils.Keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}
	}

	return m, nil
}

func (m Model) View(ctx context.Context) string {
	return helpFooterStyle.Copy().Width(ctx.ScreenWidth).Render(m.help.View(utils.Keys))
}

func (m *Model) SetWidth(width int) {
	m.help.Width = width
}
