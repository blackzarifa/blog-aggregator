package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DB_URL          string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
	configPath      string `json:"-"`
}

func Read(filename string) (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(homeDir, filename)
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	config.configPath = configPath

	return &config, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(c.configPath, data, 0600)
}
