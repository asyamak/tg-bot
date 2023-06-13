package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Token   string `json:"token"`
	TimeOut int    `json:"time_out"`
}

func NewConfig(path string) (*Config, error) {
	var config *Config
	configFile, err := os.Open(path)
	if err != nil {
		log.Printf("error to open file: %v\n", err)
	}
	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		log.Printf("error to decode configs: %v\n", err)
	}
	return &Config{
		Token:   config.Token,
		TimeOut: config.TimeOut,
	}, nil
}
