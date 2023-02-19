package config

import "fmt"

type Config struct {
	Environment string `default:"development"`
	Server      ServerConfig
	Database    DatabaseConfig
	Auth        AuthConfig
}

type ServerConfig struct {
	SERVER_PORT string `default:"8080"`
}

type DatabaseConfig struct {
	DATABASE_HOST     string `default:"localhost"`
	DATABASE_PORT     string `default:"5432"`
	DATABASE_USER     string `default:"user"`
	DATABASE_PASSWORD string `default:"password"`
	DATABASE_NAME     string `default:"db"`
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
