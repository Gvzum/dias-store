package config

// import (
// 	"fmt"
// 	"log"

// 	"github.com/Gvzum/dias-store.git/api/models"
// 	"github.com/jinzhu/gorm"
// 	"honnef.co/go/tools/config"
// )

// var DB *gorm.DB

// // Connect to the database and return a reference to the gorm.DB object
// func Connect() (*gorm.DB, error) {
// 	// Load database configuration from environment variables
// 	config.Config
// 	dbName := config.GetString("DB_NAME")
// 	dbUser := config.GetString("DB_USER")
// 	dbPassword := config.GetString("DB_PASSWORD")
// 	dbHost := config.GetString("DB_HOST")
// 	dbPort := config.GetString("DB_PORT")

// 	// Create database connection string
// 	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 		dbUser, dbPassword, dbHost, dbPort, dbName)

// 	// Connect to the database
// 	database, err := gorm.Open("mysql", dbURI)
// 	if err != nil {
// 		log.Fatalf("failed to connect to database: %s", err)
// 		return nil, err
// 	}

// 	// Set connection pool settings
// 	database.DB().SetMaxIdleConns(10)
// 	database.DB().SetMaxOpenConns(100)

// 	// Return the database connection
// 	return database, nil
// }

// // Migrate the database schema
// func Migrate(database *gorm.DB) {
// 	database.AutoMigrate(&models.User{})
// 	// Add any other models you want to migrate here
// }
