package main

import (
	"fmt"
	"sync"
	"time"
)

var balance int
var mutex sync.Mutex

func deposit(amount int) { //deposit function
	mutex.Lock()
	defer mutex.Unlock()
	balance += amount
	fmt.Printf("Deposited $%d. New balance: $%d\n", amount, balance)
}

func withdraw(amount int) { //withdraw function
	mutex.Lock()
	defer mutex.Unlock()
	if balance >= amount {
		balance -= amount
		fmt.Printf("Withdrawn $%d. New balance: $%d\n", amount, balance)
	} else {
		fmt.Println("Insufficient funds.")
	}
}

func main() {
	balance = 1000
	go deposit(200)
	go withdraw(300)
	go deposit(500)

	time.Sleep(time.Second * 10) // Allow goroutines to finish before exiting
}
