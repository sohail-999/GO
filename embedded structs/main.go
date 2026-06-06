package main

import "fmt"

type Address struct {
	City    string
	Country string
}

type Person struct {
	Name string
	Age  int
	Address
}

func main() {

	p := Person{
		Name: "Ali",
		Age:  25,
		Address: Address{
			City:    "Lahore",
			Country: "Pakistan",
		},
	}

	// Promoted access — feels like Person's own field!
	fmt.Println(p.City)    // "Lahore"
	fmt.Println(p.Country) // "Pakistan"

	// Also works the full way
	fmt.Println(p.Address)
}
