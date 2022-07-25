package db

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/woka20/sirclo-ecommerce-task/customers/src/model"
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

func GetInMemoryDb() map[string]*model.Customer {
	db := make(map[string]*model.Customer)

	for _, m := range loadCustomerFromJson() {
		cust := model.NewCustomer()
		cust.ID = m.ID
		cust.FirstName = m.FirstName
		cust.LastName = m.LastName
		cust.Email = m.Email
		cust.Password = m.Password

		birthDate, _ := time.Parse(time.RFC3339, m.BirthDate)
		cust.BirthDate = birthDate

		db[cust.ID] = cust
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
