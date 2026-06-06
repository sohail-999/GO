package main

import "fmt"

func Add[T int | float64](a, b T) T {
	return a + b
}

func Multiply[T int | float64](a, b T) T {
	return a * b
}

func Print[T any](value T) {
	fmt.Println(value)
}

func main() {
	// Type Inference - No need to specify type explicitly
	fmt.Println(Add(10, 20))     // Go infers T = int
	fmt.Println(Add(3.14, 2.86)) // Go infers T = float64

	fmt.Println(Multiply(4, 5))     // Go infers T = int
	fmt.Println(Multiply(2.5, 4.0)) // Go infers T = float64

	Print("Type Inference") // Go infers T = string
	Print(100)              // Go infers T = int
}
