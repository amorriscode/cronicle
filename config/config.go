package config

import (
	"cronicle/ui/constants"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/viper"
)

func InitConfig() {
	v := viper.GetViper()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Failed to get user's home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".config")

	currUser, err := user.Current()
	if err != nil {
		log.Fatal("Failed to get current user: %w", err)
	}

	configRoot := filepath.Join(configDir, "cronicle")
	configName := "config"
	configType := "yml"
	configPath := filepath.Join(configRoot, configName+"."+configType)

	storagePath := filepath.Join(homeDir, "cronicle")

	v.AddConfigPath(configRoot)
	v.SetConfigName(configName)
	v.SetConfigType(configType)

	// Set config defaults
	v.SetDefault(constants.CONFIG_USER, currUser.Username)
	v.SetDefault(constants.CONFIG_STORAGE_DIR, storagePath)

	// Attempt to read existing config
	if err = v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			createConfig(configRoot, configPath)
		} else {
			log.Fatal("Error reading config file %w \n")
		}
	}
}

func createConfig(configRoot string, configPath string) {
	v := viper.GetViper()

	_, err := os.Stat(configPath)
	if !os.IsExist(err) {
		if err := os.MkdirAll(configRoot, os.ModePerm); err != nil {
			log.Fatal("Failed to create cronicle config path: %w", err)
		}

		if _, err := os.Create(configPath); err != nil {
			log.Fatal("Failed to create cronicle config: %w", err)
		}
	}

	if err = v.WriteConfigAs(configPath); err != nil {
		log.Fatal(err)
	}
}
