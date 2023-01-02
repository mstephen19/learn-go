package main

import (
	"fmt"
)

type Contact struct {
	phone string
	address string
}

type Customer struct {
	name string
	contact Contact
}

func (customer Customer) sayName() {
	fmt.Println(customer.name)
}

func main() {
	customer := Customer{
		name: "Martin",
		contact: Contact{
			phone: "123-123-1234",
			address: "123 Sesame St.",
		},
	}

	customer.sayName()

	fmt.Println(customer)
}