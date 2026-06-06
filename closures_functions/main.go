package main

import "fmt"

func makeCounter() func() int {
	count := 0 // This variable is "captured" by the closure

	return func() int {
		count++ // It remembers and modifies 'count'
		return count
	}
}

func main() {
	counter := makeCounter()

	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3
}
