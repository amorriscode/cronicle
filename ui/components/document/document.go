package document

import (
	"cronicle/ui/constants"
	"cronicle/ui/context"
	"cronicle/utils"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	ctx        *context.Context
	content    string
	viewport   viewport.Model
	dimensions constants.Dimensions
}

func New(ctx *context.Context) Model {
	m := Model{
		ctx: ctx,
	}

	m.dimensions = m.getDimensions()

	m.viewport = viewport.New(m.dimensions.Width, m.dimensions.Height)

	return m
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, utils.Keys.ScrollUp):
			m.viewport.HalfViewUp()
		case key.Matches(msg, utils.Keys.ScrollDown):
			m.viewport.HalfViewDown()
		}
	}

	return m, nil
}

func (m Model) View() string {
	d, _ := glamour.Render(m.viewport.View(), "notty")
	return lipgloss.NewStyle().Render(lipgloss.JoinVertical(lipgloss.Top, d))
}

func (m *Model) getDimensions() constants.Dimensions {
	return constants.Dimensions{
		// TODO: make remove this magic number
		Height: m.ctx.ContentHeight - 6,
		Width:  m.ctx.ContentWidth,
	}
}

func (m *Model) UpdateContext(ctx *context.Context) {
	m.ctx = ctx
	m.dimensions = m.getDimensions()
	m.viewport.Width = m.dimensions.Width
	m.viewport.Height = m.dimensions.Height
}

func (m *Model) UpdateContent(content string) {
	// Empty struct used to strip frontmatter
	var f struct{}
	var c string

	// Parse frontmatter and remove it
	rest, err := frontmatter.Parse(strings.NewReader(content), &f)
	if err != nil {
		c = content
	} else {
		c = string(rest)
	}

	m.content = c
	m.viewport.SetContent(m.content)
}
