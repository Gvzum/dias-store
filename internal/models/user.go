package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string
	Email       string `gorm:"unique;unique_index;not null"`
	Password    string
	IsSuperUser bool `gorm:"default:false"`
}
