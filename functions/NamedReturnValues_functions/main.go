package main

import (
	"fmt"
)

func multiply(a int, b int) (mul int, div float64) {
	mul = a * b
	div = float64(a) / float64(b)
	return //return current values
}

func main() {
	multiple, divide := multiply(25, 4)

	fmt.Println("the multiplication :", multiple)
	fmt.Println("the division :", divide) //gives float result
}
