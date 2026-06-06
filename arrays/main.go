package main

import "fmt"

func main() {
	var a [3]string
	a[0] = "hello"
	a[1] = "s ameer "
	a[2] = "s"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	oddnumber := [6]int{3, 5, 7, 9, 11, 13}
	fmt.Println("here is the odd numbers:", oddnumber)
}
