package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	// Example configuration fields.
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

func Load(file string) (*Config, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
