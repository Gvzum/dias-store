package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `binding:"required"`
	Description string  `binding:"null"`
	Price       float64 `binding:"required"`
	ImageURL    string  `binding:"null"`
	CategoryID  uint
	Category    Category
}
