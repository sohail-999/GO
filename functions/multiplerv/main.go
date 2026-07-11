package main

import (
	"fmt"
)

func mutiplerv(a int, b int) (int, bool) {
	sum := a + b
	return sum, true
}
func main() {

	result, returnvalue := mutiplerv(5, 10)
	fmt.Println(result, returnvalue)

}
