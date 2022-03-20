package section

import (
	"cronicle/ui/components/table"
	"cronicle/ui/constants"
	"cronicle/ui/context"
	"cronicle/utils"
	"io/fs"
	"io/ioutil"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/viper"
)

type Model struct {
	ctx        *context.Context
	table      table.Model
	name       string
	storageDir string
}

func New(ctx *context.Context, name string) Model {
	v := viper.GetViper()
	sd := v.GetString("storage_dir") + "/" + name

	m := Model{
		ctx:        ctx,
		storageDir: sd,
		name:       name,
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

func (m Model) getFiles() []fs.FileInfo {
	utils.CreateDirIfNotExist(m.storageDir)

	files, err := ioutil.ReadDir(m.storageDir)
	if err != nil {
		log.Fatal(err)
	}

	return files
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
	rows := []table.Row{}

	for _, f := range m.getFiles() {
		if !f.IsDir() {
			r := table.Row{utils.FileNameWithoutExtension(f.Name()), ""}
			rows = append(rows, r)
		}
	}

	return rows
}
