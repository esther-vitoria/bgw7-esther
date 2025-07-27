package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	ApiToken string
	Database db
}

type db struct {
	Host     string
	User     string
	Password string
	Port     string
	Database string
}

type Config struct {
	Server *ServerConfig
}

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, using system environment variables")
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func NewConfig(config *Config) *Config {
	defaultConfig := &Config{
		Server: &ServerConfig{
			ApiToken: getEnv("API_KEY", ""),
			Database: db{
				Host:     getEnv("DB_HOST", ""),
				Port:     getEnv("DB_PORT", ""),
				User:     getEnv("DB_USER", ""),
				Password: getEnv("DB_PASSWORD", ""),
				Database: getEnv("DB_NAME", ""),
			},
		},
	}

	fmt.Println(*defaultConfig.Server)

	if config != nil {
		if config.Server != nil {
			defaultConfig.Server = config.Server
		}
	}

	return &Config{
		Server: defaultConfig.Server,
	}
}
