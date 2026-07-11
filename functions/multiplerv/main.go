package main

import (
	"fmt"
)

func mutiplerv(a int, b int) (int, string) {
	sum := a + b
	return sum, "numbers"
}
func main() {

	result, returnvalue := mutiplerv(5, 10)
	fmt.Println(result, returnvalue)

}
