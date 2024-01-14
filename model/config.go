package model

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	BaseURL   string   `yaml:"baseURL"`
	MaxActors int      `yaml:"maxActors"`
	Actions   []Action `yaml:"actions"`
}

func NewConfigFromYamlFile(filename string) Config {
	config := Config{}

	f, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("error reading config file: %v", err))
	}

	if err := yaml.Unmarshal(f, &config); err != nil {
		panic(fmt.Sprintf("error parsing config file: %v", err))
	}

	return config
}

func NewConfig(baseURL string, maxActors int) Config {
	return Config{
		BaseURL:   baseURL,
		MaxActors: maxActors,
		Actions:   GetSampleActions(),
	}
}
