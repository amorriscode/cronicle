package form

import (
	"cronicle/ui/components/pages"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	focusIndex int
	Inputs     []textinput.Model
	Errors     []string
	validate   func(m *Model)
	onSubmit   func(m *Model)
}

func New(inputs []textinput.Model, validate func(m *Model), onSubmit func(m *Model), focusIndex int) Model {
	return Model{
		focusIndex: focusIndex,
		Inputs:     inputs,
		validate:   validate,
		onSubmit:   onSubmit,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			// Clear out the errors
			m.Errors = []string{}

			s := msg.String()

			// Create the todo if Create button pressed
			if s == "enter" && m.focusIndex == len(m.Inputs) {
				m.validate(&m)
				if len(m.Errors) > 0 {
					return m, nil
				}

				m.onSubmit(&m)

				return m, pages.ChangePageCmd("sections")
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.Inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.Inputs)
			}

			cmds := make([]tea.Cmd, len(m.Inputs))

			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m Model) View() string {
	var b strings.Builder

	for i := range m.Inputs {
		b.WriteString(m.renderInput(i))
		if i < len(m.Inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.Inputs) {
		button = &focusedButton
	}

	if len(m.Errors) > 0 {
		b.WriteString(fmt.Sprintf("\n\n%s", strings.Join(m.Errors, "\n")))
	}

	fmt.Fprintf(&b, "\n\n%s", *button)

	return b.String()
}

func (m *Model) renderInput(input int) string {
	if m.focusIndex == input {
		m.Inputs[input].PromptStyle = focusedStyle
		m.Inputs[input].TextStyle = focusedStyle
		m.Inputs[input].Focus()
	} else {
		m.Inputs[input].PromptStyle = noStyle
		m.Inputs[input].TextStyle = noStyle
		m.Inputs[input].Blur()
	}

	return lipgloss.NewStyle().Render(
		lipgloss.JoinHorizontal(lipgloss.Top, m.Inputs[input].View()),
	)
}

func (m *Model) updateInputs(msg tea.Msg) tea.Cmd {
	var cmds = make([]tea.Cmd, len(m.Inputs))

	for i := range m.Inputs {
		m.Inputs[i], cmds[i] = m.Inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}
