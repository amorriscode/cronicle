package daily

import (
	"cronicle/ui/components/table"
	"cronicle/ui/constants"
	"cronicle/ui/context"
	"cronicle/utils"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	ctx   *context.Context
	table table.Model
}

func New(ctx *context.Context) Model {
	m := Model{
		ctx: ctx,
	}

	m.table = table.New(m.getDimensions(), m.getColumns(), m.getRows())

	return m
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	m.table, cmd = m.table.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	return m.table.View()
}

func (m *Model) UpdateContext(ctx *context.Context) {
	m.ctx = ctx
	m.table.SetDimensions(m.getDimensions())
}

func (m *Model) getDimensions() constants.Dimensions {
	return constants.Dimensions{
		Height: m.ctx.ContentHeight - 2,
		Width:  m.ctx.ContentWidth,
	}
}

func (m Model) getColumns() []table.Column {
	return []table.Column{
		{Title: "Date", Width: &dateCellWidth},
		{Title: "Tags", Grow: utils.BoolPtr(true)},
	}
}

func (m Model) getRows() []table.Row {
	return []table.Row{
		{"2022-03-01", ""},
		{"2022-03-02", "bugs, prs"},
		{"2022-03-03", "bugs, prs"},
		{"2022-03-04", ""},
		{"2022-03-05", ""},
		{"2022-03-06", ""},
		{"2022-03-07", ""},
		{"2022-03-08", ""},
		{"2022-03-09", "bugs, prs"},
		{"2022-03-10", ""},
		{"2022-03-11", ""},
		{"2022-03-12", ""},
		{"2022-03-13", ""},
		{"2022-03-14", ""},
		{"2022-03-15", ""},
		{"2022-03-16", ""},
		{"2022-03-17", "bugs, prs"},
		{"2022-03-18", ""},
		{"2022-03-19", ""},
		{"2022-03-20", ""},
		{"2022-03-21", ""},
	}
}
