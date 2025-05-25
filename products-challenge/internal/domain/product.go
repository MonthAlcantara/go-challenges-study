package domain

import "github.com/google/uuid"

type Product struct {
	ID       string `json:"id_product"`
	Price    float64
	Quantity int
}

func NewProduct(Name string, Price float64, Quantity int) *Product {
	return &Product{
		ID:       uuid.New().String(),
		Price:    Price,
		Quantity: Quantity,
	}
}
