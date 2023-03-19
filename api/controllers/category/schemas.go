package category

type CreateCategorySchema struct {
	Name string `json:"name" binding:"required"`
}

type DetailedCategorySchema struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
