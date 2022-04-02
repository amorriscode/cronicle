package help

import (
	"cronicle/ui/context"
	"cronicle/utils"

	bubblesHelp "github.com/charmbracelet/bubbles/help"
)

type Model struct {
	help bubblesHelp.Model
}

func New() Model {
	help := bubblesHelp.New()

	help.ShowAll = true

	return Model{
		help: help,
	}
}

func (m Model) View(ctx context.Context) string {
	return helpFooterStyle.Copy().Width(ctx.ScreenWidth).Render(m.help.View(utils.Keys))
}

func (m *Model) SetWidth(width int) {
	m.help.Width = width
}
