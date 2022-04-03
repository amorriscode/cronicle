package todo

import (
	"cronicle/ui/constants"
	"cronicle/ui/context"
	"cronicle/utils"
	"cronicle/utils/todo"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Column struct {
	Width *int
	Grow  *bool
}

type SectionModel struct {
	ctx      *context.Context
	todos    []string
	viewport viewport.Model
	currRow  int
}

func NewSectionUI(ctx *context.Context) SectionModel {
	m := SectionModel{
		ctx:     ctx,
		currRow: 0,
	}

	m.viewport = viewport.New(m.getDimensions().Width, m.getDimensions().Height)

	m.setTodos(getTodos())

	return m
}

func (m SectionModel) Update(msg tea.Msg) (SectionModel, tea.Cmd) {
	var cmd tea.Cmd

	if len(m.todos) > 0 {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, utils.Keys.Down):
				m.currRow = (m.currRow + 1) % len(m.todos)
			case key.Matches(msg, utils.Keys.Up):
				newRow := m.currRow - 1
				if newRow < 0 {
					newRow = len(m.todos) - 1
				}
				m.currRow = newRow
			}
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)

	return m, cmd
}

func (m SectionModel) View() string {
	m.setTodos(m.todos)

	return lipgloss.JoinVertical(lipgloss.Left, m.viewport.View())
}

func (m *SectionModel) setTodos(todos []string) {
	m.todos = todos

	renderedTodos := make([]string, 0, len(m.todos))

	for i := range m.todos {
		renderedTodos = append(renderedTodos, m.renderTodo(i))
	}

	m.viewport.SetContent(lipgloss.JoinVertical(lipgloss.Left, renderedTodos...))
}

func (m *SectionModel) renderTodo(row int) string {
	var style lipgloss.Style

	if m.currRow == row {
		style = selectedCellStyle
	} else {
		style = cellStyle
	}

	return rowStyle.Copy().Render(
		lipgloss.JoinHorizontal(lipgloss.Top, style.Copy().Render(m.todos[row])),
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

func getTodos() []string {
	var r []string
	todos := todo.GetAllTodos()
	return append(r, todos...)
}
