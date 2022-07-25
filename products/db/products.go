package db

import (
	"encoding/json"

	"github.com/woka20/sirclo-ecommerce-task/products/src/model"
)

var productData = []byte(`[{
  "id": 1,
  "name": "Zeta Gundam",
  "description": "transformation enabled figure",
  "stock": 25,
  "price": 347000
}, {
  "id": 2,
  "name": "Wing Gundam",
  "description": "Main Character Gundam",
  "stock": 47,
  "price": 310000
}, {
  "id": 3,
  "name": "Blazing Gundam",
  "description": "find the red blaze of figure",
  "stock": 22,
  "price": 280000
}, {
  "id": 4,
  "name": "Barbatos Gundam",
  "description": "new release gundam",
  "stock": 60,
  "price": 377000
}, {
  "id": 5,
  "name": "Deatschyte Gundam",
  "description": "Black armor gundam",
  "stock": 10,
  "price": 200000
}]`)

// GetProductInMemoryDb return *model.Product map, this fake database just for testing purposes only
func GetProductInDb() map[int]*model.Product {
	db := make(map[int]*model.Product)

	for _, p := range loadProductFromJson() {

		product := model.NewProduct()
		product.ID = p.ID
		product.Name = p.Name
		product.Description = p.Description
		product.Stock = p.Stock
		product.Price = p.Price

		db[product.ID] = product
	}

	return db
}

func loadProductFromJson() model.Products {
	var products model.Products
	json.Unmarshal(productData, &products)
	return products
}
