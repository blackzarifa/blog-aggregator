package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DB_URL          string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	cfg := Config{}

	fullPath, err := getConfigFilePath()
	if err != nil {
		return cfg, err
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return cfg, err

	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return cfg, err

	}

	return cfg, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return write(*cfg)
}

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}
