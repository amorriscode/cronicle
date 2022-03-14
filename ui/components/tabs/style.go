package tabs

import "github.com/charmbracelet/lipgloss"

var (
	tabsBorderHeight = 1

	tabsContentHeight = 2

	TabsHeight = tabsBorderHeight + tabsContentHeight

	tabStyle = lipgloss.NewStyle().Faint(true).Padding(0, 2)

	activeTabStyle = tabStyle.Copy().Faint(false).Bold(true)

	tabRowStyle = lipgloss.NewStyle().
			Height(tabsContentHeight).
			PaddingTop(1).
			PaddingBottom(0).
			BorderBottom(true).
			BorderStyle(lipgloss.ThickBorder())
)
