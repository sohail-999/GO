// @write a program to read the marks of the 5 subjects and then print the total and average
// give the 5 numbers
// first add them and print value
// and the total divided by 5\
package main

import "fmt"

func totalandavg(a, b, c, d, e int) {

	total, average := fmt.Println("total:", a+b+c+d+e, "average:", (a+b+c+d+e)/5)
	fmt.Println(total, average)
}

func main() {
	totalandavg(42, 45, 60, 45, 58)
}
