package main

import (
	"fmt"
	"github.com/Gvzum/dias-store.git/config/database"
	"log"

	"github.com/Gvzum/dias-store.git/internal/models"
)

func main() {
	db := database.GetDB()

	// Auto-migrate the models
	if err := db.Migrator().DropTable(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.Order{},
		&models.ProductRate{},
	); err != nil {
		log.Fatalf("Failed when back migrating %s", err)
	}

	fmt.Println("Tables droped.")
}
