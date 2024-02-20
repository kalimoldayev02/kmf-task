package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
)

var (
	instance *Config
	once     sync.Once
)

type Database struct {
	Driver   string `json:"driver"`
	DBName   string `json:"db_name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type HttpServer struct {
	Port int `json:"port"`
}

type ServiceHosts struct {
	NotionalBank string `json:"nationalbank"`
}

type Config struct {
	Database     Database     `json:"database"`
	HttpServer   HttpServer   `json:"http_server"`
	ServiceHosts ServiceHosts `json:"service_hosts"`
}

func NewCoifig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error load .env file: %s", err)
	}

	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		return nil, fmt.Errorf("CONFIG_FILE is not set in .env")
	}

	configPath := filepath.Join("configs", configFile)

	configContent, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("can not read config: %s", err)
	}

	config := &Config{}
	if err := json.Unmarshal(configContent, &config); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %s", err)
	}

	return config, nil
}
