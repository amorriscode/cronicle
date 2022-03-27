package section

import (
	"cronicle/ui/components/document"
	"cronicle/ui/components/table"
	"cronicle/ui/constants"
	"cronicle/ui/context"
	"cronicle/utils"
	"io/fs"
	"io/ioutil"
	"log"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
)

type Model struct {
	ctx             *context.Context
	table           table.Model
	document        document.Model
	documentContent string
	name            string
	storageDir      string
	files           []fs.FileInfo
	currFile        int
}

func New(ctx *context.Context, name string) Model {
	v := viper.GetViper()
	sd := v.GetString("storage_dir") + "/" + name

	m := Model{
		ctx:        ctx,
		storageDir: sd,
		name:       name,
		document:   document.New(ctx),
		currFile:   0,
	}

	m.files = m.getFiles()

	m.table = table.New(m.getDimensions(), m.getColumns(), m.getRows())

	return m
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var (
		cmds        []tea.Cmd
		documentCmd tea.Cmd
		tableCmd    tea.Cmd
	)

	if len(m.files) > 0 {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, utils.Keys.Down):
				m.currFile = (m.currFile + 1) % len(m.files)
			case key.Matches(msg, utils.Keys.Up):
				newFile := m.currFile - 1
				if newFile < 0 {
					newFile = len(m.files) - 1
				}
				m.currFile = newFile
			}
		}

		m.documentContent = m.getFile(m.storageDir + "/" + m.files[m.currFile].Name())
		m.document.UpdateContent(m.documentContent)
		m.document, documentCmd = m.document.Update(msg)
	}

	m.table, tableCmd = m.table.Update(msg)

	cmds = append(cmds, documentCmd, tableCmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return lipgloss.NewStyle().Render(lipgloss.JoinHorizontal(lipgloss.Left, m.table.View(), m.document.View()))
}

func (m Model) getFiles() []fs.FileInfo {
	utils.CreateDirIfNotExist(m.storageDir)

	files, err := ioutil.ReadDir(m.storageDir)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func (m Model) getFile(file string) string {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	return string(contents)
}

func (m *Model) UpdateContext(ctx *context.Context) {
	m.ctx = ctx
	m.table.SetDimensions(m.getDimensions())
	m.document.UpdateContext(ctx)
}

func (m *Model) getDimensions() constants.Dimensions {
	return constants.Dimensions{
		Height: m.ctx.ContentHeight - 2,
		Width:  SectionWidth,
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

	for _, f := range m.files {
		if !f.IsDir() {
			r := table.Row{utils.FileNameWithoutExtension(f.Name()), ""}
			rows = append(rows, r)
		}
	}

	return rows
}
