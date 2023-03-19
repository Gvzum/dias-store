package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"unique;unique_index;not null"`
}
