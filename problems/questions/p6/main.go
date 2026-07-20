//Write a program o swap two numbers using a third variable
//first  make a function receiving two numbers
//retunrn swap numbers
//call it

package main

import "fmt"

func swap(a int, b int) (int, int) {
	return b, a
}

func main() {

	fmt.Println(swap(111, 999))

}
