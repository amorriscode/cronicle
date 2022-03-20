package utils

import (
	"cronicle/ui/constants"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

type WriteParams struct {
	Message, Date, Tags string
}

func WriteToFile(m []string) {
	// load storage directory from config
	d := GetStorageDir()
	fn := filepath.Join(d, uuid.NewString()+".txt")

	fullMessage := strings.Join(m[:], "\n")
	log.Println("message before writing to file", fullMessage)
	log.Println("file path", fn)
	f, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(constants.ERROR_OPEN_FILE, err)
	}

	if _, err := f.Write([]byte(fullMessage)); err != nil {
		f.Close()
		log.Fatal(constants.ERROR_WRITE_FILE, err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(constants.ERROR_CLOSE_FILE, err)
	}

}

func ComposeTodo(w WriteParams) []string {
	output := make([]string, 6)
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
		output = append(output, "tags: ")
		tagArray := strings.Split(w.Tags, ",")
		for index := 0; index < len(tagArray); index++ {
			tag := fmt.Sprintf("- %s", tagArray[index])
			output = append(output, tag)
		}
	}

	// footer
	output = append(output, "---\n")

	//todo item
	message := fmt.Sprintf("- [ ] %s", w.Message)
	output = append(output, message)

	return output
}
