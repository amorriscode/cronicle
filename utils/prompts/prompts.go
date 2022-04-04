package prompts

import (
	"cronicle/ui/constants"
	"cronicle/utils"
	"cronicle/utils/types"
	"fmt"
	"io/fs"
	"log"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/manifoldco/promptui"
)

func SelectTodo(todoOptions []types.TodoProperties, action string) (int, error) {

	templates := &promptui.SelectTemplates{
		Label:    "{{ . | bold }}",
		Active:   "\U00002192  {{ .Todo | cyan }}",
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
		Active:   "\U00002192 {{ .FileName | cyan }}",
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

func GetEntryDisplayOptions(t string, files []fs.FileInfo) []types.EntryProperties {
	var matter types.EntryProperties
	var options []types.EntryProperties

	for _, f := range files {
		path := utils.GetPath([]string{t, f.Name()})
		content := utils.GetDataFromFile(path)
		rest, err := frontmatter.Parse(strings.NewReader(content), &matter)
		if err != nil {
			log.Println(err)
		}
		a := strings.Split(string(rest), "\n")
		b := strings.Join(a, ",")
		matter.Entry = utils.TruncateText(string(rest)[6:], constants.MaxLengthDisplayOption)
		matter.EntryDetails = utils.TruncateText(b, constants.MaxLengthDetails)
		matter.FileName = f.Name()
		options = append(options, matter)
	}

	return options
}

func GetTodoDisplayOptions(files []fs.FileInfo) []types.TodoProperties {
	var matter types.TodoProperties
	var options []types.TodoProperties

	for _, f := range files {
		path := utils.GetPath([]string{"todo", f.Name()})
		content := utils.GetDataFromFile(path)
		rest, err := frontmatter.Parse(strings.NewReader(content), &matter)
		if err != nil {
			log.Println(err)
		}
		matter.Todo = utils.TruncateText(string(rest)[6:], constants.MaxLengthDisplayOption)
		matter.TodoDetails = utils.TruncateText(string(rest)[6:], constants.MaxLengthDetails)
		options = append(options, matter)
	}

	return options
}
