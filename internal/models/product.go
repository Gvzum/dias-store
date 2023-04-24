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
	UserID      uint
	User        User
}

type ProductRate struct {
	gorm.Model
	Rate      float64 `binding:"required"`
	UserID    uint    `binding:"required"`
	User      User
	ProductID uint `binding:"required"`
	Product   Product
}

type Comment struct {
	gorm.Model
	Message   string `binding:"required"`
	UserID    uint   `binding:"required"`
	User      User
	ProductID uint `binding:"required"`
	Product   Product
}
