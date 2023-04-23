package product

import "github.com/go-playground/validator/v10"

type CreateProductSchema struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	ImageURL    string  `json:"image_url"`
	CategoryID  uint    `json:"category_id" binding:"required"`
}

type DetailedProductSchema struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	ImageURL     string  `json:"image_url"`
	CategoryName string  `json:"category_name"`
	UserID       uint    `json:"user_id"`
	Rate         float64 `json:"rate"`
}

type ListProductSchema struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type RateProductSchema struct {
	ProductID uint    `json:"product_id"`
	Rate      float64 `json:"rate" binding:"required,gte=0,lte=10"`
}

func (r *RateProductSchema) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
