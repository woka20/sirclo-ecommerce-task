package usecase

import (
	"github.com/woka20/sirclo-ecommerce-task/customers/src/model"
)

// UseCaseResult model
type UseCaseResult struct {
	Result interface{}
	Error  error
}

// MemberUseCase interface abstraction
type CustomerUseCase interface {
	Save(m *model.Customer) <-chan error
	FindByID(id string) <-chan UseCaseResult
	FindByEmail(email string) <-chan UseCaseResult
}
