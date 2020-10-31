package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	BindAddr         string `yaml:"bind_addr"`
	DatabaseURL      string `yaml:"database_url"`
	DbName           string `yaml:"db_name"`
	DbUserCollection string `yaml:"db_user_collection"`
}

func NewConfig(path string) (*Config, error) {
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = yaml.Unmarshal(configFile, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
