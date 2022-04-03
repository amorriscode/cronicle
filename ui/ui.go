package ui

import (
	"cronicle/ui/components/brag"
	"cronicle/ui/components/daily"
	"cronicle/ui/components/help"
	"cronicle/ui/components/pages"
	"cronicle/ui/components/section"
	"cronicle/ui/components/sections"
	"cronicle/ui/components/tabs"
	"cronicle/ui/components/todo"
	"cronicle/ui/context"
	"cronicle/utils"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	ctx          context.Context
	page         string
	sections     sections.Model
	todoCreateUI todo.CreateModel
}

func New() Model {
	m := Model{
		page:         "sections",
		sections:     sections.New(),
		todoCreateUI: todo.NewCreateUI(),
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd           tea.Cmd
		cmds          []tea.Cmd
		sectionsCmd   tea.Cmd
		todoCreateCmd tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, utils.Keys.Todo):
			m.page = "createTodo"
			return m, nil

		case key.Matches(msg, utils.Keys.Daily):
			daily.Edit()
			return m, tea.Quit

		case key.Matches(msg, utils.Keys.Brag):
			brag.Edit()
			return m, tea.Quit

		case key.Matches(msg, utils.Keys.Escape):
			m.page = "sections"

		case key.Matches(msg, utils.Keys.Quit):
			cmd = tea.Quit
		}

	case tea.WindowSizeMsg:
		m.onWindowSizeChanged(msg)

	case pages.ChangePageMsg:
		m.page = msg.Page
	}

	m.syncContext()

	if m.page == "sections" {
		m.sections, sectionsCmd = m.sections.Update(msg)
		cmds = append(cmds, sectionsCmd)
	}

	if m.page == "createTodo" {
		m.todoCreateUI, todoCreateCmd = m.todoCreateUI.Update(msg)
		cmds = append(cmds, todoCreateCmd)
	}

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.page == "createTodo" {
		return m.todoCreateUI.View()
	}

	return m.sections.View()
}

func (m *Model) onWindowSizeChanged(msg tea.WindowSizeMsg) {
	m.ctx.ScreenHeight = msg.Height
	m.ctx.ScreenWidth = msg.Width

	m.ctx.ContentHeight = msg.Height - tabs.TabsHeight - help.FooterHeight
	m.ctx.ContentWidth = msg.Width - section.SectionWidth
}

func (m *Model) syncContext() {
	m.sections.UpdateContext(&m.ctx)
}
