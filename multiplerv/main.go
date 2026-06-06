package main

import (
	"fmt"
)

func mutiplerv(a int, b int) (int, int) {
	sum := a + b
	return sum, 0
}
func main() {

	result, returnvalue := mutiplerv(5, 10)
	fmt.Println(result, returnvalue)

}
