package main

import (
	"fmt"
)

type Shark struct {
	Name string
}

func (s *Shark) Swim() {
	fmt.Println(s.Name, "is swimming")
}

func main() {
	//arr := []int{1, 2, 3, 4, 5} //array of 5 integers
	//fmt.Println("Array:", arr[len(arr)]) //length of the array
	s := &Shark{Name: "sammy"} //pointer to the shark
	s = nil                    //pointer to nothing
	s.Swim()
}
