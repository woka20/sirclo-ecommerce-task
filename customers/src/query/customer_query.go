package query

import (
	"errors"

	"github.com/woka20/sirclo-ecommerce-task/customers/src/model"
)

// memberQueryInMemory model
type customerQueryInMemory struct {
	db map[string]*model.Customer
}

// NewMemberQueryInMemory for initialise memberQueryInMemory model
func NewCustomerQueryInMemory(db map[string]*model.Customer) CustomerQuery {
	return &customerQueryInMemory{db}
}

// FindByEmail will return Member by its email
func (q *customerQueryInMemory) FindByEmail(email string) <-chan QueryResult {
	output := make(chan QueryResult)
	go func() {
		defer close(output)

		var cust *model.Customer
		for _, v := range q.db {
			if v.Email == email {
				cust = v
				break
			} else {
				output <- QueryResult{Error: errors.New("MEMBER_NOT_FOUND")}
				return
			}
		}

		output <- QueryResult{Result: cust}
	}()
	return output
}
