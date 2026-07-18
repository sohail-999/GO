// @write a program which reads the length and breadth of a rectangle
// and then prints area and perimter
// reads the length and breadth
// apply the area  formula ,area=length * width
// apply the perimeter formula,perimter=2(lenght +  width)
package main

import "fmt"

func Areaandperimeter(l int, b int) {
	fmt.Println("area:\n", l*b)
	fmt.Println("Perimeter:\n", 2*(l+b))

}
func main() {
	Areaandperimeter(2, 6)
}
