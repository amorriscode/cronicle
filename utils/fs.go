package utils

import (
	"log"
	"os"
	"path/filepath"
)

func CreateDirIfNotExist(d string) {
	_, err := os.Stat(d)
	if !os.IsExist(err) {
		if err := os.MkdirAll(d, os.ModePerm); err != nil {
			log.Fatal("Failed to create directory: %w", err)
		}
	}
}

func FileNameWithoutExtension(f string) string {
	return f[:len(f)-len(filepath.Ext(f))]
}
