package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
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
		DATABASE_HOST:     os.Getenv("DATABASE_HOST"),
		DATABASE_PORT:     os.Getenv("DATABASE_PORT"),
		DATABASE_USER:     os.Getenv("DATABASE_USER"),
		DATABASE_PASSWORD: os.Getenv("DATABASE_PASSWORD"),
		DATABASE_NAME:     os.Getenv("DATABASE_NAME"),
	}
	ServerConfig := &ServerConfig{
		SERVER_PORT:      os.Getenv("SERVER_PORT"),
		TOKEN_SECRET_KEY: os.Getenv("TOKEN_SECRET_KEY"),
	}
	AppConfig = &Config{
		Database: *DatabaseConfig,
		Server:   *ServerConfig,
	}

	return nil
}
