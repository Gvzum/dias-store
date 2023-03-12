package config

import "fmt"

type Config struct {
	Environment string `default:"development"`
	Server      ServerConfig
	Database    DatabaseConfig
	Auth        AuthConfig
}

type ServerConfig struct {
	SERVER_PORT      string
	TOKEN_SECRET_KEY string
}

type DatabaseConfig struct {
	DATABASE_HOST     string
	DATABASE_PORT     string
	DATABASE_USER     string
	DATABASE_PASSWORD string
	DATABASE_NAME     string
}

func (d *DatabaseConfig) DatabaseUrl() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		d.DATABASE_HOST,
		d.DATABASE_PORT,
		d.DATABASE_USER,
		d.DATABASE_NAME,
		d.DATABASE_PASSWORD,
	)
}

type AuthConfig struct {
	SecretKey string `default:"secret"`
	Issuer    string `default:"your-app"`
}
