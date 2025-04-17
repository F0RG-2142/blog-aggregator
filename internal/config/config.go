package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, fmt.Errorf("could not find home directory")
	}
	configFilePath := filepath.Join(homeDir, ".gatorconfig.json")
	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, fmt.Errorf("could not read config file")
	}
	var currentConfig Config
	if err := json.Unmarshal(configFile, &currentConfig); err != nil {
		return Config{}, fmt.Errorf("could not unmarshal config file")
	}
	return currentConfig, nil
}

func SetUser(user string) error {
	config, err := Read()
	if err != nil {
		return fmt.Errorf("could not read config file")
	}
	config.CurrentUserName = user
	if user == config.CurrentUserName {
		return nil
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not find home directory: %v", err)
	}
	configFilePath := filepath.Join(homeDir, ".gatorconfig.json")
	configData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal updated config file: %v", err)
	}
	if err := os.WriteFile(configFilePath, configData, 0o644); err != nil {
		return fmt.Errorf("could not write config file: %v", err)
	}
	return nil
}
