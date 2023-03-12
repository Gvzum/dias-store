package main

import (
	"fmt"
	"github.com/Gvzum/dias-store.git/config/database"
	"github.com/Gvzum/dias-store.git/internal/models"
	"log"
)

func main() {
	db := database.GetDB()

	// Auto-migrate the models
	if err := db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.Order{},
	); err != nil {
		log.Fatalf("Failed when migrating %s", err)
	}

	fmt.Println("Tables created.")
}
