package main

import (
	"fmt"
	"sync"
	"time"
)
//avoiding memory leak helps write good code.

var a string

var mu sync.Mutex

func main() {

	go func() {

		time.Sleep(100 * time.Millisecond)

		mu.Lock()

		fmt.Println("Goroutine:", a)

		mu.Unlock()

	}()

	mu.Lock()

	a = "hello, world"

	mu.Unlock()

	time.Sleep(200 * time.Millisecond)
}
