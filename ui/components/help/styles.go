package help

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	FooterHeight = 3

	helpFooterStyle = lipgloss.NewStyle().
			Height(FooterHeight - 1).
			BorderTop(true).
			BorderStyle(lipgloss.NormalBorder())
)
