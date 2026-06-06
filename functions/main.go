package main

import "fmt"

func Salamdua(name string) string {

	return "Aoa , " + name

}

func Add(a, b int) int {

	return a + b

}

// No return value? Omit the return type

func main() {

	fmt.Println(Salamdua("sahil"))
	fmt.Println(Add(10, 20))

}
