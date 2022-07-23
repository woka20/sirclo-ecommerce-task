package query

import (
	"errors"
	"sirclo-ecommerce-task/products/src/model"
)

type productQuery struct {
	db map[int]*model.Product
}

func NewProductQuery(db map[int]*model.Product) ProductQuery {
	return &productQuery{db}
}

// FindByID will return Product by its id
func (pq *productQuery) FindByID(id int) <-chan QueryResult {
	output := make(chan QueryResult)
	go func() {
		defer close(output)

		product, ok := pq.db[id]
		if !ok {
			output <- QueryResult{Error: errors.New("product not found")}
			return
		}

		output <- QueryResult{Result: product}
	}()
	return output
}

// FindAll will return all products
func (pq *productQuery) FindAll() <-chan QueryResult {
	output := make(chan QueryResult)
	go func() {
		defer close(output)

		var products model.Products
		for _, v := range pq.db {
			products = append(products, *v)
		}

		output <- QueryResult{Result: products}
	}()
	return output
}
