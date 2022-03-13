package tabs

import "github.com/charmbracelet/lipgloss"

var (
	tabsBorderHeight  = 1
	tabsContentHeight = 2
	TabsHeight        = tabsBorderHeight + tabsContentHeight

	tab = lipgloss.NewStyle().Faint(true).Padding(0, 2)

	active = tab.Copy().Faint(false).Bold(true)

	tabRow = lipgloss.NewStyle().
		Height(tabsContentHeight).
		PaddingTop(1).
		PaddingBottom(0).
		BorderBottom(true).
		BorderStyle(lipgloss.ThickBorder())
)
