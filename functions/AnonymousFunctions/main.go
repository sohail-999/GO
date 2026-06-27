package main

import "fmt"

func main() {
	counter := 0
	increment := func() {

		counter++

		fmt.Println(counter)
	}
	increment() //prints 1
	increment() //prints 2
	increment() //prints 3
	increment() //prints 4
}
