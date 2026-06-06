package main

import "fmt"

func double(n int) int {
	n = n * 2 // modifies the local copy only
	return n
}

func main() {

	x := 5

	double(x)

	fmt.Println(x) // still 5
}
