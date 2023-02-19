package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var AppConfig *Config

func LoadEnv() error {
	envFile := ".env"

	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("failed to load env file: %s", envFile)
		return err
	}

	DatabaseConfig := &DatabaseConfig{
		DATABASE_HOST:     GetEnv("DATABASE_HOST", ""),
		DATABASE_PORT:     GetEnv("DATABASE_PORT", ""),
		DATABASE_USER:     GetEnv("DATABASE_USER", ""),
		DATABASE_PASSWORD: GetEnv("DATABASE_PASSWORD", ""),
		DATABASE_NAME:     GetEnv("DATABASE_NAME", ""),
	}
	ServerConfig := &ServerConfig{
		SERVER_PORT: GetEnv("SERVER_PORT", ""),
	}
	AppConfig = &Config{
		Database: *DatabaseConfig,
		Server:   *ServerConfig,
	}

	return nil
}

func GetEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
