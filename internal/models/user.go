package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `gorm:"type:varchar(50)"`
	Email       string `gorm:"unique;unique_index;not null"`
	Password    string
	IsSuperUser bool `gorm:"default:false"`
}
