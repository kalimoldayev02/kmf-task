package config

import (
	"encoding/json"
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

func GetInstance() *Config {
	if instance == nil {
		once.Do(func() { // принимает в аргумент функцию, которая отработает один раз за вызов
			instance = loadConfig()
		})
	}

	return instance
}

func loadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error load .env file: %s", err)
	}

	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		log.Fatalf("CONFIG_FILE is not set in .env")
	}

	configPath := filepath.Join("configs", configFile)

	configContent, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("can not read config: %s", err)
	}

	var config Config
	if err := json.Unmarshal(configContent, &config); err != nil {
		log.Fatalf("error decodingJSON : %s", err)
	}

	return &config
}
