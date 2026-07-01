package main

import "fmt"

func main() {
	//ch := make(chan int)
	//ch <- 43
	//value := <-ch

	ch := make(chan int) //creating an unbufferred channel
	go func() {
		fmt.Println("sending the 42 to the channel:")
		ch <- 42
		fmt.Println("sent 42")
	}()
	fmt.Println("receiving from the channel:")
	value := <-ch
	fmt.Println("value received", value)
}
