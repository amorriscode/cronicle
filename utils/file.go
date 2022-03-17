package utils

import (
	"cronicle/ui/constants"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type WriteParams struct {
	message, date, tags string
}

func WriteToFile(m string) {
	// load storage directory from config
	d := GetStorageDir()
	fn := filepath.Join(d, uuid.NewString()+".txt")
	if w.date == "" {
		// due date set to end of month if not set
		t := time.Date(time.Now().Year(), time.Now().Month(), 29, 0, 0, 0, 0, time.UTC)
		w.date = t.Format("2006-01-02")
	}

	fullMessage := w.message + " " + w.date

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
	output := make([]string, 3)
	// date created
	t := time.Now()
	formatedTime := t.Format("2006-01-02 15:04")
	date := fmt.Sprintf("date: %s \n", formatedTime)
	// date due
	if w.date != "" {
		dueDate := fmt.Sprintf("due: %s \n", w.date)
		output = append(output, dueDate)
	}
	// tags
	if w.tags != "" {
		for i, c := range w.tags {
			if c
		}
	}

	output = append(output, date, "type: todo \n")

	return output
}
