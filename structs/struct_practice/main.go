package main //defining the main package

import (
	"fmt" //importing the fmt package
)

type employeedata struct { //creating a struct named employeedata
	Name       string
	Email      string
	streetno   int
	Age        int
	IsEmployed bool
}

func main() {

	Alice := employeedata{ //making a variable of Alice and assigning the values to it

		Name:       "Alice",
		Email:      "alice@example.com",
		streetno:   8,
		Age:        30,
		IsEmployed: false,
	}

	fmt.Println(Alice.Name, Alice.Email, Alice.streetno, Alice.Age, Alice.IsEmployed) //for printing the data of Alice
}
