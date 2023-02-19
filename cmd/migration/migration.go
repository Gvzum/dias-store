package main

import (
	"fmt"

	"github.com/Gvzum/dias-store.git/cmd/server/config"
	"github.com/Gvzum/dias-store.git/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		fmt.Printf("Failed to load env file: %s", err)
	}

	db, err := gorm.Open(postgres.Open(config.AppConfig.Database.DatabaseUrl()), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// Auto-migrate the models
	db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.Order{},
	)

	fmt.Println("Tables created.")
}
