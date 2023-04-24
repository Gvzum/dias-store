package product

type CreateProductSchema struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	ImageURL    string  `json:"image_url"`
	CategoryID  uint    `json:"category_id" binding:"required"`
}

type ListProductSchema struct {
	ID            uint    `json:"id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	AverageRating float64 `json:"average_rating"`
}

type DetailedProductSchema struct {
	ListProductSchema
	Description  string  `json:"description"`
	ImageURL     string  `json:"image_url"`
	CategoryName string  `json:"category_name"`
	UserID       uint    `json:"user_id"`
	Rate         float64 `json:"rate"`
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
