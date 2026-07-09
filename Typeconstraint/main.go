package main

import "fmt"

func Print[T any](value T) { //T is a type parameter that can be any type
	fmt.Println(value)
}

func Swap[T any](a, b T) (T, T) { //t any accept any type and return the same type
	return b, a
}

func main() {

	Print(42)
	Print("Hello")
	Print(3.14)

	a, b := Swap(7, 6)
	fmt.Println(a, b) //Output: 6 7

	x, y := Swap("Go", "Language")
	fmt.Println(x, y) //Output: Language Go
	m, n := Swap("seveeen", "sixxxx")
	fmt.Println(m, n)
}
