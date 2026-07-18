package main

import "fmt"

func Areaandperimeter(length int, breadth int) {
	fmt.Println("area :", length*breadth)
	fmt.Println("Perimeter:", 2*(length+breadth))

}
func main() {
	Areaandperimeter(2, 6)
}
