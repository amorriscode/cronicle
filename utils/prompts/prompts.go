package prompts

import (
	"cronicle/utils/types"
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

func SelectTodo(todoOptions []types.TodoProperties, action string) (int, error) {

	templates := &promptui.SelectTemplates{
		Label:    "{{ . | bold }}",
		Active:   "\U00002705  {{ .Todo | cyan }}",
		Inactive: "{{ .Todo | cyan }}",
		Selected: "\U00002714  {{ .Todo | green | cyan }}",
		Details: `
		--------- Properties ----------
		{{ "Date:" | faint }}	{{ .Date }}
		{{ "Tags:" | faint }}	{{ .Tags }}
		{{ "Due:" | faint }}	{{ .Due }}
		{{ "Todo:" | faint }}	{{ .TodoDetails }}`,
	}

	searcher := func(input string, index int) bool {
		todo := todoOptions[index]
		name := strings.Replace(strings.ToLower(todo.Todo), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     fmt.Sprintf("Select todo to %s:", action),
		Items:     todoOptions,
		Templates: templates,
		Searcher:  searcher,
	}

	id, _, err := prompt.Run()

	return id, err
}

func SelectEntry(options []types.EntryProperties, action string) (int, error) {

	templates := &promptui.SelectTemplates{
		Label:    "{{ . | bold }}",
		Active:   "\U00002705  {{ .FileName | cyan }}",
		Inactive: "{{ .FileName | cyan }}",
		Selected: "\U00002714  {{ .FileName | green | cyan }}",
		Details: `
		--------- Properties ----------
		{{ "Date:" | faint }}	{{ .Date }}
		{{ "Tags:" | faint }}	{{ .Tags }}
		{{ "Entry:" | faint }}	{{ .EntryDetails }}`,
	}

	searcher := func(input string, index int) bool {
		entry := options[index]
		name := strings.Replace(strings.ToLower(entry.Entry), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     fmt.Sprintf("Select entry to %s:", action),
		Items:     options,
		Templates: templates,
		Searcher:  searcher,
	}

	id, _, err := prompt.Run()

	return id, err
}
