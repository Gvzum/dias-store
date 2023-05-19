package order

type ItemSchema struct {
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type CreateOrderSchema struct {
	OrderItems []ItemSchema `json:"items"`
}
