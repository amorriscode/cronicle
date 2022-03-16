package utils

import (
	"cronicle/ui/constants"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

func WriteToFile(m string, date time.Time) {
	// load storage directory from config
	d := GetStorageDir()
	fn := filepath.Join(d, uuid.NewString()+".txt")
	if date == "" {
		// due date set to end of month if not set
		date = time.Date(time.Now().Year(), time.Now().Month(), 29, 0, 0, 0, 0, time.UTC)
	}
	fullMessage := m + " " + date

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
