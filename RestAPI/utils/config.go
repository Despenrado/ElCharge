package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config struct stored api parameters
type Config struct {
	BindAddr         string `yaml:"bind_addr"`
	DatabaseURL      string `yaml:"database_url"`
	DbName           string `yaml:"db_name"`
	DbUserCollection string `yaml:"db_user_collection"`
	RedisDB          string `yaml:"db_redis"`
	JWTKey           string `yaml:"jwtKey"`
}

// NewConfig Constructor
// Loads params from 'yaml' file
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
