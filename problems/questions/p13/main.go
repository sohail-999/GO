// Write a program to read the three numbers and then find the smallest among them
// first get the three  numbers from the users as the input
// apply the some condition to to find the smallest number
package main

import "fmt"

func main() {

	var a, b, c int
	fmt.Println("Please enter the numbers:")

	fmt.Scan(&a, &b, &c)

	smallest := a
	if b < smallest {
		smallest = b
	}
	if c < smallest {
		smallest = c
	}
	fmt.Println("Smallest number:", smallest)
}
