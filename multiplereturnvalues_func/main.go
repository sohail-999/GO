package main

import "fmt"

func add(a int, b int) (int, int) { //multiple return types
	// returns TWO ints
	sum := a + b

	product := a * b

	return sum, product
}

func main() {

	s, p := add(6, 6)

	fmt.Println(s) // 12

	fmt.Println(p) // 36

}
