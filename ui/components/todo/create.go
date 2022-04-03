package todo

import (
	"cronicle/ui/components/pages"
	"cronicle/utils"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type CreateModel struct {
	focusIndex int
	inputs     []textinput.Model
	errors     []string
}

func NewCreateUI() CreateModel {
	m := CreateModel{
		focusIndex: 0,
		inputs:     make([]textinput.Model, 3),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Task"
			t.Focus()
		case 1:
			t.Placeholder = "Comma-separated tags (optional)"
		case 2:
			t.Placeholder = "Due date - YYYY-MM-DD (optional)"
		}

		m.inputs[i] = t
	}

	return m
}

func (m CreateModel) Update(msg tea.Msg) (CreateModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			// Clear out the errors
			m.errors = []string{}

			s := msg.String()

			// Create the todo if Create button pressed
			if s == "enter" && m.focusIndex == len(m.inputs) {
				m.validateInput()
				if len(m.errors) > 0 {
					return m, nil
				}

				m.createTodo()

				return m, pages.ChangePageCmd("sections")
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}

				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m CreateModel) View() string {
	var b strings.Builder

	b.WriteString("Create a todo\n")

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}

	if len(m.errors) > 0 {
		b.WriteString(fmt.Sprintf("\n\n%s", strings.Join(m.errors, "\n")))
	}

	fmt.Fprintf(&b, "\n\n%s", *button)

	return b.String()
}

func (m *CreateModel) updateInputs(msg tea.Msg) tea.Cmd {
	var cmds = make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m *CreateModel) validateInput() {
	if m.inputs[0].Value() == "" {
		m.errors = append(m.errors, "task cannot be empty")
	}
}

func (m CreateModel) createTodo() {
	todo := utils.ComposeTodo(utils.WriteParams{Message: m.inputs[0].Value(), Tags: m.inputs[1].Value(), Date: m.inputs[2].Value()})
	utils.WriteToFile(todo, utils.GetPath([]string{"todo", uuid.NewString() + ".md"}))
}
