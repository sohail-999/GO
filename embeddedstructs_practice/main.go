package main

import (
	"fmt"
)

type employeedata struct {
	Name      string
	Email     string
	Age       int
	Ismarried bool
	address   employeeaddress
}

type employeeaddress struct {
	Street       string
	Streetnumber int
}

func main() {

	p := employeedata{
		Name:      "Alice",
		Email:     "alice@example.com",
		Age:       30,
		Ismarried: false,

		address: employeeaddress{
			Street:       "Main Street",
			Streetnumber: 88,
		},
	}
	fmt.Println("Employee Data:", p.Name, p.Email, p.Age, p.Ismarried)
	fmt.Println("Employee Address:", p.address.Street, p.address.Streetnumber)

}
