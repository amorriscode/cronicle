package todo

import (
	"cronicle/ui/constants"
	"cronicle/ui/context"
	"cronicle/utils"
	"log"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Column struct {
	Width *int
	Grow  *bool
}

type Row []string

type SectionModel struct {
	ctx      *context.Context
	Rows     []Row
	viewport viewport.Model
	currRow  int
}

func NewSectionUI(ctx *context.Context) SectionModel {
	// todos := utils.GetAllTodos()

	m := SectionModel{
		ctx:     ctx,
		currRow: 0,
	}

	m.viewport = viewport.New(m.getDimensions().Width, m.getDimensions().Height)

	rows := getTodos()
	m.SetRows(rows)

	return m
}

func (m SectionModel) Update(msg tea.Msg) (SectionModel, tea.Cmd) {
	var cmd tea.Cmd

	if len(m.Rows) > 0 {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, utils.Keys.Down):
				m.currRow = (m.currRow + 1) % len(m.Rows)
			case key.Matches(msg, utils.Keys.Up):
				newRow := m.currRow - 1
				if newRow < 0 {
					newRow = len(m.Rows) - 1
				}
				m.currRow = newRow
			}
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)

	return m, cmd
}

func (m SectionModel) View() string {
	m.SetRows(m.Rows)

	return lipgloss.JoinVertical(lipgloss.Left, m.viewport.View())
}

func (m *SectionModel) SetRows(rows []Row) {
	m.Rows = rows

	renderedRows := make([]string, 0, len(m.Rows))

	for i := range m.Rows {
		renderedRows = append(renderedRows, m.renderRow(i))
	}

	m.viewport.SetContent(lipgloss.JoinVertical(lipgloss.Left, renderedRows...))
}

func (m *SectionModel) renderRow(row int) string {
	var style lipgloss.Style

	if m.currRow == row {
		style = selectedCellStyle
	} else {
		style = cellStyle
	}

	renderedColumns := make([]string, 2)

	for _, column := range m.Rows[row] {
		renderedColumns = append(
			renderedColumns,
			style.Copy().Render(column),
		)
	}

	return rowStyle.Copy().Render(
		lipgloss.JoinHorizontal(lipgloss.Top, renderedColumns...),
	)
}

func (m *SectionModel) getDimensions() constants.Dimensions {
	return constants.Dimensions{
		Height: m.ctx.ContentHeight - 2,
		Width:  m.ctx.ContentWidth,
	}
}

func (m *SectionModel) UpdateContext(ctx *context.Context) {
	m.ctx = ctx
	m.viewport.Height = m.getDimensions().Height
	m.viewport.Width = m.getDimensions().Width
}

func getTodos() []Row {
	var todoRows []Row

	todos := utils.GetAllTodos()

	for _, todoPath := range todos {
		todo := utils.GetTodoFromFile(todoPath)
		log.Println(todo)
		// todoRows = append(todoRows, Row{todo.Text})
	}

	todoRows = append(todoRows, Row{"Todo 1", "Todo 2", "Todo 3"})

	return todoRows
}
