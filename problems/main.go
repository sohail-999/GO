package main

import (
	"fmt"
	"sync"
)

type Foo struct {
	firstDone  chan struct{}
	secondDone chan struct{}
}

func NewFoo() *Foo {
	return &Foo{
		firstDone:  make(chan struct{}),
		secondDone: make(chan struct{}),
	}
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()

	close(f.firstDone) //closing
}

func (f *Foo) Second(printSecond func()) {
	<-f.firstDone

	// Do not change this line
	printSecond()

	close(f.secondDone)
}

func (f *Foo) Third(printThird func()) {
	<-f.secondDone

	// Do not change this line
	printThird()
}

func main() {
	foo := NewFoo()

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		foo.First(func() {
			fmt.Print("first")
		})
	}()

	go func() {
		defer wg.Done()
		foo.Second(func() {
			fmt.Print("second")
		})
	}()

	go func() {
		defer wg.Done()
		foo.Third(func() {
			fmt.Print("third")
		})
	}()

	wg.Wait()
}
