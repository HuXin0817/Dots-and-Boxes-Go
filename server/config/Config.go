package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ListenOn string `yaml:"ListenOn"`
}

func NewFromFile(file string) (conf *Config, err error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(data, &conf); err != nil {
		return nil, err
	}
	return conf, nil
}
