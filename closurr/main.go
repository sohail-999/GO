package main

import "fmt"

func main() {
	// Create a multiplier factory
	multiply := makeMultiplier(5)

	fmt.Println(multiply(10)) // Output: 50
	fmt.Println(multiply(20)) // Output: 100

	// Create another multiplier with different factor
	triple := makeMultiplier(3)
	fmt.Println(triple(10)) // Output: 30
}

func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor // Closure captures 'factor'
	}
}
