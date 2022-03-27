package utils

import (
	"cronicle/ui/constants"
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

type UpdateParams struct {
	Message, Date, Tags string
	Number              int
}

func WriteToFile(m []string, t string) {
	// load storage directory from config
	d := GetStorageDir()
	CreateDirIfNotExist(filepath.Join(d, t))
	fn := ""

	if t == "todo" {
		fn = GetPath([]string{t, uuid.NewString() + ".md"})
	} else {
		time := time.Now()
		date := time.Format("2006-01-02")
		fn = GetPath([]string{t, date + ".md"})
	}

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
