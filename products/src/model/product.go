package model

import (
	"time"
)

// Product
type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stock       int       `json:"stock"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

//
// Products type
type Products []Product

// NewProduct function for initialise Product model
func NewProduct() *Product {
	now := time.Now()
	return &Product{
		CreatedAt: now,
		UpdatedAt: now,
	}
}
