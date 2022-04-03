package todo

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	cellStyle = lipgloss.NewStyle().
			PaddingLeft(1).
			PaddingRight(1).
			MaxHeight(1)

	selectedCellStyle = cellStyle.Copy().
				Foreground(lipgloss.Color("#000000")).
				Background(lipgloss.Color("#FFFFFF"))

	rowStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true)

	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	noStyle      = lipgloss.NewStyle()

	focusedButton = focusedStyle.Copy().Render("[ Create ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Create"))
)
