package table

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	headerHeight = 2

	cellStyle = lipgloss.NewStyle().
			PaddingLeft(1).
			PaddingRight(1).
			MaxHeight(1)

	selectedCellStyle = cellStyle.Copy().
				Background(lipgloss.Color("#FFFFFF"))

	titleCellStyle = cellStyle.Copy().
			Bold(true)

	singleRuneTitleCellStyle = titleCellStyle.Copy().Width(1)

	headerStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true)

	rowStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true)
)
