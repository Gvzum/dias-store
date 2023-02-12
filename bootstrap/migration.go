package bootstrap

// import (
// 	"log"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func MigrateDatabase() {
// 	db, err := gorm.Open(postgres.Open("./database.db"), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	if err = db.AutoMigrate(); err != nil {
// 		log.Println(err)
// 	}
// }
