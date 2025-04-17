package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Could not find home directory")
	}
	configFilePath := homeDir + "/.gatorconfig.json"
	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		panic("Could not read config file: " + err.Error())
	}
	var currentConfig Config
	json.Unmarshal(configFile, &currentConfig)
	return currentConfig
}
