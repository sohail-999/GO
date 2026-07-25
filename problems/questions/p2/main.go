// write a program to read two numbers and print their sum
// first the user gives two numbers
// second add them
// then print them
package main

import "fmt"

func sum(u int, v int) int {
	return u + v

}

func main() {
	var print int
	print = sum(9, 8)
	fmt.Println(print)
}
