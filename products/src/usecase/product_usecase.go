package usecase

import (
	"errors"

	"github.com/woka20/sirclo-ecommerce-task/products/src/model"
	"github.com/woka20/sirclo-ecommerce-task/products/src/query"
)

// productUseCaseImpl model
type productUseCaseImpl struct {
	productQuery query.ProductQuery
}

// NewProductUseCase for initialise productUseCaseImpl model
func NewProductUseCase(productQuery query.ProductQuery) ProductUseCase {
	return &productUseCaseImpl{
		productQuery: productQuery,
	}
}

// FindByID
func (u *productUseCaseImpl) FindByID(id int) <-chan UseCaseResult {
	output := make(chan UseCaseResult)

	go func() {
		productResult := <-u.productQuery.FindByID(id)

		if productResult.Error != nil {
			output <- UseCaseResult{Error: productResult.Error}
			return
		}

		product, ok := productResult.Result.(*model.Product)

		if !ok {
			output <- UseCaseResult{Error: errors.New("Result is not Product")}
			return
		}

		output <- UseCaseResult{Result: product}
	}()

	return output
}

// FindAll
func (u *productUseCaseImpl) FindAll() <-chan UseCaseResult {
	output := make(chan UseCaseResult)

	go func() {
		productResult := <-u.productQuery.FindAll()

		if productResult.Error != nil {
			output <- UseCaseResult{Error: productResult.Error}
			return
		}

		products, ok := productResult.Result.(model.Products)

		if !ok {
			output <- UseCaseResult{Error: errors.New("Result is not Products")}
			return
		}

		output <- UseCaseResult{Result: products}
	}()

	return output
}
