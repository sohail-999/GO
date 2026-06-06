package main

import (
	"fmt"
)

func main() {

	var counter int = 0

	increement := func() {
		counter++
		fmt.Println("Counter:", counter)

	}
	increement()
	increement()
	increement()
	increement()

}
