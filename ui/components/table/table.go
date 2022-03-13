package table

import (
	"cronicle/ui/constants"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Column struct {
	Title string
	Width *int
	Grow  *bool
}

type Row []string

type Model struct {
	Columns    []Column
	Rows       []Row
	dimensions constants.Dimensions
	viewport   viewport.Model
}

func New(dimensions constants.Dimensions, columns []Column, rows []Row) Model {
	m := Model{
		Columns:    columns,
		Rows:       rows,
		dimensions: dimensions,
		viewport:   viewport.New(dimensions.Width, dimensions.Height),
	}

	m.SetRows(rows)

	return m
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	m.viewport, cmd = m.viewport.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	header := m.renderHeader()
	body := m.renderBody()

	return lipgloss.JoinVertical(lipgloss.Left, header, body)
}

func (m *Model) renderHeaderColumns() []string {
	renderedColumns := make([]string, len(m.Columns))
	takenWidth := 0
	numGrowingColumns := 0

	for i, column := range m.Columns {
		if column.Grow != nil && *column.Grow {
			numGrowingColumns += 1
			continue
		}

		if column.Width != nil {
			renderedColumns[i] = titleCellStyle.Copy().
				Width(*column.Width).
				MaxWidth(*column.Width).
				Render(column.Title)
			takenWidth += *column.Width
			continue
		}

		cell := titleCellStyle.Copy().Render(column.Title)
		renderedColumns[i] = cell
		takenWidth += lipgloss.Width(cell)
	}

	leftoverWidth := m.dimensions.Width - takenWidth
	if numGrowingColumns == 0 {
		return renderedColumns
	}

	growCellWidth := leftoverWidth / numGrowingColumns
	for i, column := range m.Columns {
		if column.Grow == nil || !*column.Grow {
			continue
		}

		renderedColumns[i] = titleCellStyle.Copy().Width(growCellWidth).MaxWidth(growCellWidth).Render(column.Title)
	}

	return renderedColumns
}

func (m *Model) SetRows(rows []Row) {
	m.Rows = rows

	headerColumns := m.renderHeaderColumns()
	renderedRows := make([]string, 0, len(m.Rows))
	for i := range m.Rows {
		renderedRows = append(renderedRows, m.renderRow(i, headerColumns))
	}

	m.viewport.SetContent(lipgloss.JoinVertical(lipgloss.Left, renderedRows...))
}

func (m *Model) renderRow(row int, headerColumns []string) string {
	var style lipgloss.Style
	// if m.rowsViewport.GetCurrItem() == rowId {
	// 	style = selectedCellStyle
	// } else {
	// 	style = cellStyle
	// }

	renderedColumns := make([]string, len(m.Columns))
	for i, column := range m.Rows[row] {
		colWidth := lipgloss.Width(headerColumns[i])
		renderedColumns = append(
			renderedColumns,
			style.Copy().Width(colWidth).MaxWidth(colWidth).Render(column),
		)
	}

	return rowStyle.Copy().Render(
		lipgloss.JoinHorizontal(lipgloss.Top, renderedColumns...),
	)
}

func (m *Model) renderHeader() string {
	headerColumns := m.renderHeaderColumns()
	header := lipgloss.JoinHorizontal(lipgloss.Top, headerColumns...)
	return headerStyle.Copy().Width(m.dimensions.Width).MaxWidth(m.dimensions.Width).Render(header)
}

func (m *Model) renderBody() string {
	return m.viewport.View()
}
