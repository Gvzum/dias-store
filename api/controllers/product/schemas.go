package product

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

type BaseCommentSchema struct {
	Message string `json:"message" binding:"required"`
}

type UserCommentSchema struct {
	BaseCommentSchema
	ID     string `json:"id" binding:"require"`
	UserID uint   `json:"user_id" binding:"required"`
}

type CommentProductSchema struct {
	BaseCommentSchema
	ProductID uint `json:"product_id"`
	//Message   string `json:"message" binding:"required"`
}
