package ui

import (
	"cronicle/ui/components/help"
	"cronicle/ui/components/section"
	"cronicle/ui/components/sections"
	"cronicle/ui/components/tabs"
	"cronicle/ui/context"

	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	ctx      context.Context
	sections sections.Model
}

func New() Model {
	m := Model{
		sections: sections.New(),
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd         tea.Cmd
		cmds        []tea.Cmd
		sectionsCmd tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.onWindowSizeChanged(msg)
	}

	m.syncContext()

	m.sections, sectionsCmd = m.sections.Update(msg)

	cmds = append(cmds, cmd, sectionsCmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := strings.Builder{}

	s.WriteString(m.sections.View())

	return s.String()
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
