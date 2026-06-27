package main

import "fmt"

func main() {
	counter := 0
	increment := func() {

		counter++

		fmt.Println(counter)
	}
	increment()
	increment()
	increment()
	increment()
}
