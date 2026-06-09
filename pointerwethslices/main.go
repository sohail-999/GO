package main

import (
	"fmt"
)

func slicerip(s *[]int) { //s is the slice

	for i := range *s {

		(*s)[i] *= 2
	}
}

func main() {
	s := []int{5, 15, 50}

	slicerip(&s)

	fmt.Println(s)

}
