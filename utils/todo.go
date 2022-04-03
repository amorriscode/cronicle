package utils

import (
	"cronicle/ui/constants"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ComposeTodo(w WriteParams) string {
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
	path := GetPath([]string{"todo", f.Name()})
	todo := GetDataFromFile(path)
	checkedTodo := CheckTodo(todo)
	tags := ParseHeader(todo).Tags
	// get today's log file, update log file: add to exisiting tags, append complted todo
	// if file does not exist, call create log file and update log file
	WriteOrCreateDaily(WriteDailyParams{Message: checkedTodo, Tags: strings.Join(tags, ",")})
	DeleteTodo(f.Name())
}

func DeleteTodo(fileName string) {
	dirPath := GetPath([]string{"todo"})
	e := os.Remove(filepath.Join(dirPath, fileName))

	if e != nil {
		log.Fatal(constants.ERROR_DELETE_FILE, e)
	}
}

func CheckTodo(todo string) string {
	var c strings.Builder

	m := ParseContent(todo)
	c.WriteString(m[2:3])
	c.WriteString("x")
	c.WriteString(m[4:])

	return c.String()
}

func GetTodoFilePaths() []fs.FileInfo {
	p := GetPath([]string{"todo"})

	f, err := ioutil.ReadDir(p)
	if err != nil {
		log.Fatal(constants.ERROR_LIST_FILE, err)
	}

	return f
}

func GetTodoFromFile(fileName string) string {
	path := GetPath([]string{"todo", fileName})
	todo := GetDataFromFile(path)
	return todo
}

func GetAllTodos() []string {
	var todos []string

	files := GetTodoFilePaths()

	for _, f := range files {
		todo := GetTodoFromFile(f.Name())
		todos = append(todos, ParseContent(todo))
	}

	return todos
}

func ListTodos() {
	files := GetTodoFilePaths()

	for i, f := range files {
		todo := GetTodoFromFile(f.Name())
		task := ParseContent(todo)

		fmt.Printf("%v. %s", i+1, task[6:])
	}
}
