package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	// Using new() to allocate memory for a Person struct
	p := new(Person)

	// Initializing the field
	// 
	p.Name = "John Doe"
	p.Age = 30
	p.Gender = "Male"

	fmt.Println(p)
}
