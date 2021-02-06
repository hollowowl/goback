package config

import (
	"encoding/json"
	"os"
)

type DBConfig struct {
	Type string `json:"type" binding:"required"`
	DSN  string `json:"dsn" binding:"required"`
}

type ServerConfig struct {
	Addr string `json:"addr" binding:"required"`
}


type Config struct {
	Database DBConfig `json:"db" binding:"required"`
	Server ServerConfig `json:"server" binding:"required"`
}

func FromJson(filename string) (*Config, error) {
	f, _ := os.Open(filename)
	defer f.Close()
	decoder := json.NewDecoder(f)
	config := &Config{}
	err := decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
