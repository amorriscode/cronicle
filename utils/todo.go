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

func GetTodoFromFile(f fs.FileInfo) string {
	path := GetPath([]string{"todo", f.Name()})
	dat, _ := os.ReadFile(path)
	return string(dat)
}

func GetAllTodos() []fs.FileInfo {
	dirPath := GetPath([]string{"todo"})
	files, _ := ioutil.ReadDir(dirPath)
	return files
}

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
		tags := fmt.Sprintf("tags: [%s]\n", w.Tags)
		output.WriteString(tags)
	}

	// footer
	output.WriteString("---\n")

	//todo item
	message := fmt.Sprintf("- [ ] %s\n", w.Message)
	output.WriteString(message)
	return output.String()
}

func MarkCompleted(f fs.FileInfo) {
	//add todo list  to log
	time := time.Now()
	date := time.Format("2006-01-02")
	todoArr := GetTodoFromFile(f)
	WriteToFile(todoArr, GetPath([]string{"log", date + ".md"}))
	DeleteTodo(f.Name())
}

func DeleteTodo(fileName string) {
	dirPath := GetPath([]string{"todo"})
	e := os.Remove(filepath.Join(dirPath, fileName))

	if e != nil {
		log.Fatal(constants.ERROR_DELETE_FILE, e)
	}
}
