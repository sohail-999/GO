package main

import (
	"fmt"
)

func mutiplerv(a int, b int) (int, float64) {
	sum := a + b
	return sum, 6.777777
}
func main() {

	result, returnvalue := mutiplerv(5, 10)
	fmt.Println(result, returnvalue)

}
