package utils

import (
	"cronicle/ui/constants"
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func WriteToFile(m string) {
	// load storage directory from config
	d := GetStorageDir()
	fn := filepath.Join(d, uuid.NewString()+".txt")

	f, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(constants.ERROR_OPEN_FILE, err)
	}

	if _, err := f.Write([]byte(m)); err != nil {
		f.Close()
		log.Fatal(constants.ERROR_WRITE_FILE, err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(constants.ERROR_CLOSE_FILE, err)
	}

}
