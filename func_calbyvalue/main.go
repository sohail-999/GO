package main

import (
	"fmt"
)

func callbyvlaue(a int) int {
	a = a + 10
	return a
}

func main() {

	var x int = 10
	fmt.Println("Before calling the function", x)
	callbyvlaue(x)
	fmt.Println("After calling the function", x)

}
