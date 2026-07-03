package main

import "fmt"

func main() {

	ch := make(chan int, 3)
	go func() {

		fmt.Println("gorutine : waiting the value:")
		fmt.Println("now", <-ch)

	}()

	ch <- 1
	fmt.Println("sent 1")
	ch <- 2
	fmt.Println("sent 2")
	ch <- 3
	fmt.Println("sent 3")
	fmt.Println("now the space is full")
	ch <- 4
	fmt.Println("sent 4")

}
