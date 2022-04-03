package todo

import (
	"cronicle/ui/constants"
	"cronicle/utils"
	"cronicle/utils/entries"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func ComposeTodo(w utils.WriteParams) string {
	output := strings.Builder{}
	// header
	output.WriteString("---\n")

	// date created
	t := time.Now()
	formatedTime := t.Format("2006-01-02 15:04")
	date := fmt.Sprintf("date: %s\n", formatedTime)
	output.WriteString(date)

	// date due
	if w.Date != "" {
		dueDate := fmt.Sprintf("due: %s\n", w.Date)
		output.WriteString(dueDate)
	}

	// type
	output.WriteString("type: todo\n")

	// tags
	if w.Tags != "" {
		tags := fmt.Sprintf("tags: [%s]\n", strings.TrimSpace(w.Tags))
		output.WriteString(tags)
	}

	// footer
	output.WriteString("---\n")

	//todo item
	message := fmt.Sprintf("- [ ] %s\n", strings.TrimSpace(w.Message))
	output.WriteString(message)
	return output.String()
}

func MarkCompleted(f fs.FileInfo) {
	//add todo list  to log
	path := utils.GetPath([]string{"todo", f.Name()})
	todo := utils.GetDataFromFile(path)
	checkedTodo := CheckTodo(todo)
	tags := utils.ParseHeader(todo).Tags
	//add completed todo to log
	entries.WriteOrCreateEntry(utils.WriteParams{Message: checkedTodo, Tags: strings.Join(tags, ",")}, "daily")
	utils.DeleteFile(f.Name(), "todo")
}

func CheckTodo(todo string) string {
	m := utils.ParseContent(todo)

	if strings.Contains(m, "[x]") {
		return m[2:]
	}

	var c strings.Builder
	c.WriteString(m[2:3])
	c.WriteString("x")
	c.WriteString(m[4:])

	return c.String()
}

func GetTodoFilePaths() []fs.FileInfo {
	p := utils.GetPath([]string{"todo"})

	f, err := ioutil.ReadDir(p)
	if err != nil {
		log.Fatal(constants.ERROR_LIST_FILE, err)
	}

	return f
}

func GetTodoFromFile(fileName string) string {
	path := utils.GetPath([]string{"todo", fileName})
	todo := utils.GetDataFromFile(path)
	return todo
}

func GetAllTodos() []string {
	var todos []string

	files := GetTodoFilePaths()

	for _, f := range files {
		todo := GetTodoFromFile(f.Name())
		todos = append(todos, utils.ParseContent(todo))
	}

	return todos
}

func ListTodos() {
	files := GetTodoFilePaths()

	for i, f := range files {
		todo := GetTodoFromFile(f.Name())
		task := utils.ParseContent(todo)

		fmt.Printf("%v. %s", i+1, task[6:])
	}
}
