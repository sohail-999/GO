package main

import "fmt"

// sum accepts any number of integers
func sum(nums ...int) int {
	total := 0

	for _, n := range nums {
		total += n
	}
	return total
}

// greet mixes a fixed param with a variadic one
func greet(greeting string, names ...string) {
	for _, name := range names {
		fmt.Printf("%s, %s!\n", greeting, name)
	}
}

func main() {
	// 1️⃣  Pass individual arguments
	fmt.Println(sum(1, 2, 3))        // 6
	fmt.Println(sum(10, 20, 30, 40)) // 100

	// 2️⃣  Pass zero arguments — totally valid
	fmt.Println(sum()) // 0

	// 3️⃣  Spread an existing slice using the ... operator
	numbers := []int{5, 10, 15, 15}
	fmt.Println(sum(numbers...)) // 30

	// 4️⃣  Mixed fixed + variadic params
	greet("Hello", "Alice", "Bob", "joney", "Charlie")
	// Hello, Alice!
	// Hello, Bob!
	// Hello, Charlie!
}
