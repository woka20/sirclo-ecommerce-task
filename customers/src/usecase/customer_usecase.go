package usecase

import (
	"errors"

	"github.com/woka20/sirclo-ecommerce-task/customers/src/model"
	"github.com/woka20/sirclo-ecommerce-task/customers/src/query"
	"github.com/woka20/sirclo-ecommerce-task/customers/src/repo"
)

type customerUseCaseImpl struct {
	customerRepository repo.CustomerRepository
	customerQuery      query.CustomerQuery
}

func NewCustomerUseCase(customerRepository repo.CustomerRepository, customerQuery query.CustomerQuery) CustomerUseCase {
	return &customerUseCaseImpl{
		customerRepository: customerRepository,
		customerQuery:      customerQuery,
	}
}

// Save for saving Member model
func (mu *customerUseCaseImpl) Save(m *model.Customer) <-chan error {
	output := make(chan error)

	go func() {

		defer close(output)

		memberExistResult := <-mu.customerQuery.FindByEmail(m.Email)

		if memberExistResult.Error != nil && memberExistResult.Error.Error() != "MEMBER_NOT_FOUND" {
			output <- memberExistResult.Error
			return
		} else if memberExistResult.Error != nil {
			output <- memberExistResult.Error
			return
		} else {
			err := <-mu.customerRepository.Save(m)

			if err != nil {
				output <- err
				return
			}
		}

		output <- nil

	}()

	return output
}

// FindByID for load Member by its ID
func (mu *customerUseCaseImpl) FindByID(id string) <-chan UseCaseResult {
	output := make(chan UseCaseResult)

	go func() {
		custResult := <-mu.customerRepository.Load(id)

		if custResult.Error != nil {
			output <- UseCaseResult{Error: custResult.Error}
			return
		}

		cust, ok := custResult.Result.(*model.Customer)

		if !ok {
			output <- UseCaseResult{Error: errors.New("Result is not member")}
			return
		}

		output <- UseCaseResult{Result: cust}

	}()

	return output
}

// FindByEmail for load Member by its Email
func (mu *customerUseCaseImpl) FindByEmail(email string) <-chan UseCaseResult {
	output := make(chan UseCaseResult)

	go func() {
		custResult := <-mu.customerQuery.FindByEmail(email)

		if custResult.Error != nil {
			output <- UseCaseResult{Error: custResult.Error}
			return
		}

		cust, ok := custResult.Result.(*model.Customer)

		if !ok {
			output <- UseCaseResult{Error: errors.New("Result is not member")}
			return
		}

		output <- UseCaseResult{Result: cust}

	}()

	return output
}
