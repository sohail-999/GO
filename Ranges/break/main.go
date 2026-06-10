package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {

		if i == 9 {
			fmt.Println("9 is my favorite number")
			break

		}
		fmt.Println("here is the number")
		fmt.Println(i)
	}
	fmt.Println("Exit the loop")
}
