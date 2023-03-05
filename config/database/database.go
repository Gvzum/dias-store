package database

import (
	"fmt"
	"github.com/Gvzum/dias-store.git/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func init() {
	if err := config.LoadEnv(); err != nil {
		fmt.Printf("Failed to load env file: %s", err)
	}
	fmt.Println("I am not here??")

	db, err = gorm.Open(postgres.Open(config.AppConfig.Database.DatabaseUrl()), &gorm.Config{})

	if err != nil {
		panic(err)
	}

}

func GetDB() *gorm.DB {
	return db
}
