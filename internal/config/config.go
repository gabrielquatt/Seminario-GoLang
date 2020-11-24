package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// ConfigDB ...
type ConfigDB struct {
	Type   string `yaml:"type"`
	Driver string `yaml:"driver"`
	Conn   string `yaml:"conn"`
}

// CollectionConfig ...
type CollectionConfig struct {
	Version string `yaml:"version"`
}

// Config ...
type Config struct {
	DB         ConfigDB         `yaml:"db"`
	Version    string           `yaml:"version"`
	Collection CollectionConfig `yaml:"collection"`
}

// LoadConfig ...
func LoadConfig(filename string) *Config {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var c = &Config{}
	err = yaml.Unmarshal(file, c)
	if err != nil {
		panic(err)
	}

	return c
}
