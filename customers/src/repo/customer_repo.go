package repo

import (
	"errors"
	"time"

	"sirclo-ecommerce-task/customers/src/model"
)

// memberRepositoryInMemory model
type customerRepositoryInMemory struct {
	db map[string]*model.Customer
}

// NewMemberRepositoryInMemory for initialise memberRepositoryInMemory model
func NewCustomerRepositoryInMemory(db map[string]*model.Customer) CustomerRepository {
	return &customerRepositoryInMemory{db}
}

// Save function for saving Member model
func (r *customerRepositoryInMemory) Save(m *model.Customer) <-chan error {
	output := make(chan error)
	go func() {
		defer close(output)

		customer, ok := r.db[m.ID]
		if !ok {

			r.db[m.ID] = m
			output <- nil
			return
		} else {

			customer.UpdatedAt = time.Now()
			r.db[m.ID] = customer
			output <- nil
			return
		}
	}()
	return output
}

// Load for load Member by its ID
func (r *customerRepositoryInMemory) Load(id string) <-chan RepositoryResult {
	output := make(chan RepositoryResult)
	go func() {
		defer close(output)

		customer, ok := r.db[id]
		if !ok {
			output <- RepositoryResult{Error: errors.New("member not found")}
			return
		}

		output <- RepositoryResult{Result: customer}
	}()
	return output
}
