package main

import (
	"fmt"
)

func main() {

	var array [5]int
	array = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array length:", len(array), "Capacity:", cap(array))
	fmt.Printf("array : \n%p", &array)

	slice := array[1:5]
	fmt.Println("\nOriginal slice:", slice)
	fmt.Println("Slice length:", len(slice), "Capacity:", cap(slice))
	fmt.Printf("Slice ptr: %p\n", &slice[0])

	newslice := append(slice, 100) //a new address got allocated cause the capacity exceeds and  a new slice  got allocated with DOUBLE capacity
	fmt.Println("\nAfter append:", newslice)
	fmt.Println("Newslice length:", len(newslice), "Capacity:", cap(newslice))
	fmt.Printf("Newslice ptr: %p\n", &newslice[0])

	fmt.Printf("\narray:%p\n", &array)
	fmt.Println("after appending:", newslice)

}
