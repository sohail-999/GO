package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {

		if i == 4 {
			break // stop the loop
		}
		fmt.Println(i)
	}

}
