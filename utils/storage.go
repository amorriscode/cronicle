package utils

import (
	"cronicle/ui/constants"
	"log"
	"os"

	"github.com/spf13/viper"
)

func GetStorageDir() string {
	v := viper.GetViper()
	d := v.GetString(constants.CONFIG_STORAGE_DIR)

	if d == "" {
		log.Fatal("Failed to find storage directory")
	}

	_, err := os.Stat(d)
	if !os.IsExist(err) {
		if err := os.MkdirAll(d, os.ModePerm); err != nil {
			log.Fatal("Failed to create storage directory: %w", err)
		}
	}

	return d
}
