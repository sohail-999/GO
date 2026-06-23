package main

import (
	"fmt"
)

func multiply(a int, b int) (mul int, div int) {
	mul = a * b
	div = a / b
	return
}

func main() {
	multiple, divide := multiply(25, 5)

	fmt.Println("the multiplication :", multiple)
	fmt.Println("the division :", divide)
}
