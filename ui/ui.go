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
	ctx             context.Context
	page            string
	sections        sections.Model
	todoCreateForm  todo.CreateModel
	dailyCreateForm daily.CreateModel
	bragCreateForm  brag.CreateModel
}

func New() Model {
	m := Model{
		page:            "sections",
		sections:        sections.New(),
		todoCreateForm:  todo.NewCreateForm(),
		dailyCreateForm: daily.NewCreateForm(),
		bragCreateForm:  brag.NewCreateForm(),
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd            tea.Cmd
		cmds           []tea.Cmd
		sectionsCmd    tea.Cmd
		todoCreateCmd  tea.Cmd
		dailyCreateCmd tea.Cmd
		bragCreateCmd  tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, utils.Keys.Todo):
			if m.page == "sections" {
				return m, pages.ChangePageCmd("createTodo")
			}

		case key.Matches(msg, utils.Keys.Daily):
			if m.page == "sections" {
				return m, pages.ChangePageCmd("createDaily")
			}

		case key.Matches(msg, utils.Keys.Brag):
			if m.page == "sections" {
				return m, pages.ChangePageCmd("createBrag")
			}

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
		m.todoCreateForm, todoCreateCmd = m.todoCreateForm.Update(msg)
		cmds = append(cmds, todoCreateCmd)
	}

	if m.page == "createDaily" {
		m.dailyCreateForm, dailyCreateCmd = m.dailyCreateForm.Update(msg)
		cmds = append(cmds, dailyCreateCmd)
	}

	if m.page == "createBrag" {
		m.bragCreateForm, bragCreateCmd = m.bragCreateForm.Update(msg)
		cmds = append(cmds, bragCreateCmd)
	}

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.page == "createTodo" {
		return m.todoCreateForm.View()
	}

	if m.page == "createDaily" {
		return m.dailyCreateForm.View()
	}

	if m.page == "createBrag" {
		return m.bragCreateForm.View()
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
