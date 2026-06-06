package main

import "fmt"

func main() {

	array := [5]int{1, 2, 3, 4, 5}

	slice := array[:]

	fmt.Println(slice)
	

}
