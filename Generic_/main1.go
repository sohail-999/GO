package main

import "fmt"

//@the commented func are the same fuctions with diff types but same logic

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

//@  below is the generic function for accepting 2 types int,float64 values

func maxintfloat[T int | float64](a, b T) T {

	if a > b {
		return a
	}
	return b
}

func main() {

	value := maxintfloat(6.66, 7.77)
	fmt.Println("max vlaue :\n", value)

	//@ Below the code for the using the different functions for the same logic with different types

	//intvalue := maxint(6, 7)
	//fmt.Println("max int value:\n", intvalue)
	//floatvalue := maxfloat(6.66, 7.77)
	//fmt.Println("max float value  value:\n", floatvalue)

	//var s string = ",go"
	//fmt.Println("helloworld", s)

}
