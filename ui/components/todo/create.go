package todo

import (
	"cronicle/ui/components/form"
	"cronicle/utils"
	"cronicle/utils/todo"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type CreateModel struct {
	form form.Model
}

func NewCreateForm() CreateModel {
	m := CreateModel{}

	inputs := make([]textinput.Model, 3)

	var t textinput.Model
	for i := range inputs {
		t = textinput.New()
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Comma-separated tags (optional)"
		case 1:
			t.Placeholder = "Due date - YYYY-MM-DD (optional)"
		case 2:
			t.Placeholder = "Task"
			t.Focus()
		}

		inputs[i] = t
	}

	m.form = form.New(inputs, validate, onSubmit, 2)

	return m
}

func (m CreateModel) Update(msg tea.Msg) (CreateModel, tea.Cmd) {
	var cmd tea.Cmd
	m.form, cmd = m.form.Update(msg)
	return m, cmd
}

func (m CreateModel) View() string {
	return m.form.View()
}

func validate(m *form.Model) {
	if m.Inputs[2].Value() == "" {
		m.Errors = append(m.Errors, "task cannot be empty")
	}
}

func onSubmit(m *form.Model) {
	todo := todo.ComposeTodo(utils.WriteParams{Tags: m.Inputs[0].Value(), Date: m.Inputs[1].Value(), Message: m.Inputs[2].Value()})
	utils.WriteToFile(todo, utils.GetPath([]string{"todo", uuid.NewString() + ".md"}))
}
