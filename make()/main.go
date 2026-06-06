package main

import (
	"fmt"
)

func main() {

	s := make([]int, 10, 15)

	for i := 0; i < 14; i++ {

		s = append(s, i+1)

	}

	fmt.Println(s)

}
