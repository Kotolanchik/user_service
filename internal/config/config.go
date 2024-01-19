package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type Config struct {
	AppUserServicePort int    `yaml:"app_user_service_port"`
	PostgresConnStr    string `yaml:"postgres_conn_str"`
}

func GetConfigFromYAML() (*Config, error) {
	exePath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("failed to get executable path: %v", err)
	}

	configFile, err := os.ReadFile(filepath.Join(filepath.Dir(exePath), "internal/config/config.yaml"))
	if err != nil {
		return nil, fmt.Errorf("failed to read config.yaml: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config.yaml: %v", err)
	}

	return &config, nil
}
