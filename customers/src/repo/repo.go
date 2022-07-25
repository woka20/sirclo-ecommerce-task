package repo

import (
	"github.com/woka20/sirclo-ecommerce-task/customers/src/model"
)

// RepositoryResult model
type RepositoryResult struct {
	Result interface{}
	Error  error
}

type CustomerRepository interface {
	Save(m *model.Customer) <-chan error
	Load(id string) <-chan RepositoryResult
}
