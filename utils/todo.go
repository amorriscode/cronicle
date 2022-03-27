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

func GetTodoFromFile(f fs.FileInfo) []string {
	path := GetPath([]string{"todo", f.Name()})
	dat, _ := os.ReadFile(path)
	fileArr := strings.Split(string(dat), "\n")
	return fileArr
}

func GetAllTodos() []fs.FileInfo {
	dirPath := GetPath([]string{"todo"})
	files, _ := ioutil.ReadDir(dirPath)
	return files
}

func ComposeTodo(w WriteParams) []string {
	output := make([]string, 0)
	// header
	output = append(output, "---")

	// date created
	t := time.Now()
	formatedTime := t.Format("2006-01-02 15:04")
	date := fmt.Sprintf("date: %s", formatedTime)
	output = append(output, date)

	// date due
	if w.Date != "" {
		dueDate := fmt.Sprintf("due: %s", w.Date)
		output = append(output, dueDate)
	}

	// type
	output = append(output, "type: todo")

	// tags
	if w.Tags != "" {
		tags := fmt.Sprintf("tags:[%s]", w.Tags)
		output = append(output, tags)
	}

	// footer
	output = append(output, "---\n")

	//todo item
	message := fmt.Sprintf("- [ ] %s", w.Message)
	output = append(output, message)

	return output
}

func MarkCompleted(f fs.FileInfo) {
	//add todo list  to log
	todoArr := GetTodoFromFile(f)
	WriteToFile(todoArr, "log")
	DeleteTodo(f.Name())
}

func DeleteTodo(fileName string) {
	dirPath := GetPath([]string{"todo"})
	e := os.Remove(filepath.Join(dirPath, fileName))

	if e != nil {
		log.Fatal(constants.ERROR_DELETE_FILE, e)
	}
}
