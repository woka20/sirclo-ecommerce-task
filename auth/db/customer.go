package db

import (
	"encoding/json"
	"fmt"

	"sirclo-ecommerce-task/auth/src/model"
)

var customerData = []byte(`[{
  "id": "4189642134",
  "firstName": "Joyous",
  "lastName": "Billson",
  "email": "jbillson0@pinterest.com",
  "password": "UlVHwnLZhkYY",
  "passwordSalt": "1gBXQU",
  "birthDate": "2017-05-11 19:20:52"
}, {
  "id": "2715470592",
  "firstName": "Ashlen",
  "lastName": "Dronsfield",
  "email": "adronsfield1@epa.gov",
  "password": "e4qSgq",
  "passwordSalt": "yghbdqB3WZ",
  "birthDate": "2017-09-14 19:22:26"
}]`)

func GetInMemoryDb() map[string]*model.Identity {
	db := make(map[string]*model.Identity)

	for _, m := range loadCustomerFromJson() {
		cust := new(model.Identity)
		cust.ID = m.ID
		cust.Email = m.Email
		cust.Password = m.Password

		db[cust.Email] = cust
	}

	return db
}

type Customer struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	BirthDate string `json:"birthDate"`
}

type Customers []Customer

func loadCustomerFromJson() Customers {
	var custs Customers
	err := json.Unmarshal(customerData, &custs)

	if err != nil {
		fmt.Println(err)
	}
	return custs
}
