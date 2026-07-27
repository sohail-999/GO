//Write a program to read the three numbers and finds the largest among them
//first read the three numbers
//apply the condition to find the largest number

package main

import "fmt"

func main() {

	var a, b, c int
	fmt.Println("please enter the number:")
	fmt.Scan(&a, &b, &c)
	largest := a
	if b > largest {
		largest = b
	}
	if c > largest {
		largest = c
	}
	fmt.Println("the largest number is", largest)
}
