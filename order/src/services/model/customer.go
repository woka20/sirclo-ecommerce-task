package model

import (
	"time"
)

type Customer struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	BirthDate time.Time `json:"birthDate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Customers []Customer

// NewCustomer function for initialise Customer model
func NewCustomer() *Customer {
	now := time.Now()
	return &Customer{
		CreatedAt: now,
		UpdatedAt: now,
	}
}
