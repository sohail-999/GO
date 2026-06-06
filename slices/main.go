package main

import "fmt"

func main() {
	//primes := [6]int{2, 3, 5, 7, 11, 13}

	//var s []int = primes[1:3]
	//fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	//slices

	fruits := []string{"Apple", "Banana", "Cherry"}
	fmt.Println(fruits)
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4] // Creates a slice from index 1 to 3
	fmt.Println(slice)
}
