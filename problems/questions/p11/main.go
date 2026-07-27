// Write a program to read a number whether its positive or negative or zero.
// first get the number
// apply the condition to check
package main

import "fmt"

func main() {
	var number int
	fmt.Println("Please enter the number:")
	fmt.Scan(&number)
	if number > 0 {
		fmt.Println("The number is positive")
	} else {
		fmt.Println("The number is negative")
	}

}
