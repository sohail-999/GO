package main

import "fmt"

func IsEqual[T comparable](a, b T) bool {
	return a == b
}

func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {

		if v == item {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(IsEqual(10, 10))
	fmt.Println(IsEqual("Go", "Java"))

	fmt.Println(Contains([]int{1, 2, 3, 4}, 5))
	fmt.Println(Contains([]string{"Go", "Java", "Python"}, "Rust"))
}
