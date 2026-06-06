package main

import "fmt"

func main() {

	number := 1000

	fmt.Printf("Here is your number i think :")

	switch {

	case number == 0:

		fmt.Println("number is 0")

	case number > 1:

		fmt.Println("number is greater than 1")

	default:

		fmt.Println("number is less than 0")
	}

}
