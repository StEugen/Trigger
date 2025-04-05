package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Trigger struct {
	Event      string `yaml:"event"`
	Identifier string `yaml:"identifier"`
	Action     string `yaml:"action"`
}

type Config struct {
	TriggerEnvironment string    `yaml:"trigger_environment"`
	Triggers           []Trigger `yaml:"triggers"`
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

