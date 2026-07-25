// Write a program to read a number and check whether its even or odd
// first get the number
// apply some condition
package main

import "fmt"

func main() {
	var a uint
	fmt.Scan(&a)
	if a%2 == 0 {

		fmt.Println("the numebr is even")
	} else {

		fmt.Println("the number is odd")
	}
}
