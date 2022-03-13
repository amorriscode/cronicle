package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	v := viper.GetViper()
	log.Println("inside config")
	//set default user name to whoami
	v.SetDefault("user_name", "$USER")
	v.SetDefault("file_location", "$HOME/cronicle")

	// name + type of config file
	v.SetConfigName("cronicle")
	v.SetConfigType("yaml")
	// path to look for config file
	v.AddConfigPath(".")
	configDir := os.Getenv("XDG_CONFIG_HOME")
	if configDir == "" {
		configDir = "$HOME/.config"
	}
	v.AddConfigPath(configDir + "/cronicle/")
	v.AddConfigPath("/etc/cronicle/")
	log.Println(configDir, v.GetString("user_name"), v.GetString("file_location"))
	// Find and read the config file and assigns err variable
	err := v.ReadInConfig()

	if err != nil {
		log.Panic("Error reading config file %w \n")
	}
}
