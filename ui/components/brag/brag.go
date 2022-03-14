package brag

import (
	"cronicle/ui/components/table"
	"cronicle/ui/constants"
	"cronicle/ui/context"
	"cronicle/utils"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	ctx   *context.Context
	table table.Model
}

func New(ctx *context.Context) Model {

	log.Println(ctx.ScreenWidth)
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
		{"2022-01-01", ""},
		{"2022-01-02", "infra"},
		{"2022-01-03", ""},
		{"2022-01-04", ""},
		{"2022-01-05", "promotion"},
		{"2022-01-06", ""},
		{"2022-01-07", ""},
		{"2022-01-08", ""},
		{"2022-01-09", ""},
		{"2022-01-10", "prs"},
		{"2022-01-11", ""},
		{"2022-01-12", ""},
		{"2022-01-13", ""},
		{"2022-01-14", "hackathon"},
		{"2022-01-15", ""},
		{"2022-01-16", ""},
		{"2022-01-17", ""},
		{"2022-01-18", ""},
		{"2022-01-19", "mentoring, go"},
		{"2022-01-20", ""},
		{"2022-01-21", ""},
	}
}
