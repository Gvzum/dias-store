package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
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
		DATABASE_HOST:     viper.GetString("DATABASE_HOST"),
		DATABASE_PORT:     viper.GetString("DATABASE_PORT"),
		DATABASE_USER:     viper.GetString("DATABASE_USER"),
		DATABASE_PASSWORD: viper.GetString("DATABASE_PASSWORD"),
		DATABASE_NAME:     viper.GetString("DATABASE_NAME"),
	}
	ServerConfig := &ServerConfig{
		SERVER_PORT:      viper.GetString("SERVER_PORT"),
		TOKEN_SECRET_KEY: viper.GetString("TOKEN_SECRET_KEY"),
	}
	AppConfig = &Config{
		Database: *DatabaseConfig,
		Server:   *ServerConfig,
	}

	return nil
}

//func viper.GetString(key string) string {
//	//if value, exists := os.LookupEnv(key); exists {
//	//	return value
//	//}
//	value, _ := os.LookupEnv(key)
//
//	return value
//}
