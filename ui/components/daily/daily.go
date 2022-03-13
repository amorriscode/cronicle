package daily

import (
	"cronicle/ui/components/table"
	"cronicle/ui/constants"
	"cronicle/utils"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	table table.Model
}

func New() Model {
	dimensions := constants.Dimensions{Width: 20, Height: 20}

	m := Model{}

	m.table = table.New(dimensions, m.getColumns(), m.getRows())

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

func (m Model) getColumns() []table.Column {
	return []table.Column{
		{Title: "Date", Grow: utils.BoolPtr(true)},
	}
}

func (m Model) getRows() []table.Row {
	return []table.Row{
		{"2022-03-01"},
		{"2022-03-02"},
		{"2022-03-03"},
		{"2022-03-04"},
		{"2022-03-05"},
		{"2022-03-06"},
		{"2022-03-07"},
		{"2022-03-08"},
		{"2022-03-09"},
		{"2022-03-10"},
		{"2022-03-11"},
		{"2022-03-12"},
		{"2022-03-13"},
		{"2022-03-14"},
		{"2022-03-15"},
		{"2022-03-16"},
		{"2022-03-17"},
		{"2022-03-18"},
		{"2022-03-19"},
		{"2022-03-20"},
		{"2022-03-21"},
	}
}
