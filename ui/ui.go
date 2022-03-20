package ui

import (
	"cronicle/ui/components/brag"
	"cronicle/ui/components/daily"
	"cronicle/ui/components/help"
	"cronicle/ui/components/section"
	"cronicle/ui/components/tabs"
	"cronicle/ui/constants"
	"cronicle/ui/context"
	"cronicle/utils"

	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	currSection int
	keys        utils.KeyMap
	ctx         context.Context
	daily       section.Model
	brag        section.Model
	tabs        tabs.Model
	help        help.Model
}

func New() Model {
	m := Model{
		currSection: 0,
		keys:        utils.Keys,
		tabs:        tabs.New(),
		help:        help.New(),
	}

	// TODO: abstract sections out to be more dynamic
	m.daily = daily.New(&m.ctx)
	m.brag = brag.New(&m.ctx)

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd      tea.Cmd
		cmds     []tea.Cmd
		helpCmd  tea.Cmd
		dailyCmd tea.Cmd
		bragCmd  tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.PrevSection):
			prevSection := m.getPrevSection()
			m.setCurrSection(prevSection)

		case key.Matches(msg, m.keys.NextSection):
			nextSection := m.getNextSection()
			m.setCurrSection(nextSection)

		case key.Matches(msg, m.keys.Quit):
			cmd = tea.Quit
		}

	case tea.WindowSizeMsg:
		m.onWindowSizeChanged(msg)
	}

	m.syncContext()

	m.help, helpCmd = m.help.Update(msg)
	m.daily, dailyCmd = m.daily.Update(msg)
	m.brag, bragCmd = m.brag.Update(msg)

	cmds = append(cmds, cmd, helpCmd, dailyCmd, bragCmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := strings.Builder{}

	s.WriteString(m.tabs.View(m.ctx))

	s.WriteString("\n")

	s.WriteString(m.getCurrSection().View())

	s.WriteString("\n")

	s.WriteString(m.help.View(m.ctx))

	return s.String()
}

func (m *Model) setCurrSection(newSection int) {
	m.currSection = newSection
	m.tabs.SetCurrSection(newSection)
}

func (m Model) getNextSection() int {
	return (m.currSection + 1) % len(constants.Sections)
}

func (m Model) getPrevSection() int {
	m.currSection = (m.currSection - 1) % len(constants.Sections)

	if m.currSection < 0 {
		m.currSection += len(constants.Sections)
	}

	return m.currSection
}

type Section interface {
	View() string
}

func (m *Model) getCurrSection() Section {
	if m.currSection == 1 {
		return m.brag
	}

	return m.daily
}

func (m *Model) onWindowSizeChanged(msg tea.WindowSizeMsg) {
	m.help.SetWidth(msg.Width)

	m.ctx.ScreenHeight = msg.Height
	m.ctx.ScreenWidth = msg.Width

	m.ctx.ContentHeight = msg.Height - tabs.TabsHeight - help.FooterHeight
	m.ctx.ContentWidth = msg.Width
}

func (m *Model) syncContext() {
	m.daily.UpdateContext(&m.ctx)
	m.brag.UpdateContext(&m.ctx)
}
