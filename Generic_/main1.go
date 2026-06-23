package main

import "fmt"

/*(func maxint(a int, b int) int {

	if a > b {
		return a
	}
	return b

}
func maxfloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b

}*/

func maxintfloat[T int | float64](a, b T) T {

	if a > b {
		return a
	}
	return b
}

func main() {

	value := maxintfloat(6.66, 7.77)
	fmt.Println("max vlaue :\n", value)

	//intvalue := maxint(6, 7)
	//fmt.Println("max int value:\n", intvalue)
	//floatvalue := maxfloat(6.66, 7.77)
	//fmt.Println("max float value  value:\n", floatvalue)

	//var s string = ",go"
	//fmt.Println("helloworld", s)

}
