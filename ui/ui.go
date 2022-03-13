package ui

import (
	"cronicle/ui/components/tabs"
	"cronicle/ui/constants"
	"cronicle/utils"

	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	currSection int
	tabs        tabs.Model
	keys        utils.KeyMap
}

func NewModel() Model {
	return Model{
		currSection: 0,
		tabs:        tabs.NewModel(),
		keys:        utils.Keys,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

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
	}

	return m, cmd
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

func (m Model) View() string {
	s := strings.Builder{}
	s.WriteString(m.tabs.View())
	return s.String()
}
