// @write a program that reads the celcius and then convert it to the farenheit
// provide the celius value
// apply the formula on the farenheit formula
package main

import "fmt"

func celciusintofarenheit(c int) {

	fmt.Println("celcius into farenheit:", ((c * 9 / 5) + 32))

}

func main() {

	celciusintofarenheit(6)

}
