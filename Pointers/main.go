package main

import "fmt"

func main() {

	i, j := 42, 2700

	point := &i 
	// point to i
	fmt.Println(*point)
	// read i through the pointer
	*point = 21 
	// set i through the pointer
	fmt.Println(i)

	fmt.Println(i, j) // see the new value of i

	point = &j
	// point to j
	*point = *point / 27
	// divide j through the pointer
	fmt.Println(j) // see the new value of j 

}
