package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {

		if i == 9 {

			fmt.Println("\n my favorite number is hidden ")
			fmt.Println("\nlet's continue to the next number")
			fmt.Println("\n")
			continue

		}
		fmt.Println("here is the number")
		fmt.Println(i)
	}
	fmt.Println("Exit the loop")
}
