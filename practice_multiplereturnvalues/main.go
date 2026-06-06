package main

import (
	"fmt"
)

func multiplereturn(a int, b int) (int, string) {
	sum := a + b
	return sum, "the sum is : "
}

func main() {
	result, message := multiplereturn(10, 20)
	fmt.Println(message, result)

}
