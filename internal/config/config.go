package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = "gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	// Read the json file from HOME directory
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	// Decode the josn string into a new Config struct
	content, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	configStruct := Config{}
	err = json.Unmarshal(content, &configStruct)
	if err != nil {
		return Config{}, err
	}

	return configStruct, nil
}

func (config *Config) SetUser(userName string) error {
	// Convert userName string to josn format
	config.CurrentUserName = userName

	return write(*config)
}

// Some helper functions
func getConfigFilePath() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	configFilePath := filepath.Join(currentDir, configFileName)

	return configFilePath, nil
}

func write(cfg Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	// 0644 file mode represents owner can read and write, others can read
	err = os.WriteFile(configFilePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
