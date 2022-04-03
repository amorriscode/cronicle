package brag

import (
	"cronicle/ui/components/form"
	"cronicle/utils"
	"cronicle/utils/entries"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type CreateModel struct {
	form form.Model
}

func NewCreateForm() CreateModel {
	m := CreateModel{}

	inputs := make([]textinput.Model, 2)

	var t textinput.Model
	for i := range inputs {
		t = textinput.New()
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Comma-separated tags (optional)"
		case 1:
			t.Placeholder = "Entry"
			t.Focus()
		}

		inputs[i] = t
	}

	m.form = form.New(inputs, validate, onSubmit, 1)

	return m
}

func (m CreateModel) Update(msg tea.Msg) (CreateModel, tea.Cmd) {
	var cmd tea.Cmd
	m.form, cmd = m.form.Update(msg)
	return m, cmd
}

func (m CreateModel) View() string {
	return lipgloss.NewStyle().Render(lipgloss.JoinVertical(lipgloss.Top, lipgloss.NewStyle().Render("Create a brag entry\n"), m.form.View()))
}

func validate(m *form.Model) {
	if m.Inputs[1].Value() == "" {
		m.Errors = append(m.Errors, "entry cannot be empty")
	}
}

func onSubmit(m *form.Model) {
	log.Println("HEYYY SUBMIT BRAG")
	entries.WriteOrCreateEntry(utils.WriteParams{Tags: m.Inputs[0].Value(), Message: m.Inputs[1].Value()}, "brag")
}
