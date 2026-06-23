package main

import (
	"fmt"
)

func printslice[T any](items []T) {

	for _, item := range items {

		fmt.Println(item)

	}

}
func main() {

	names := []bool{ //here we specify the type
		true, false}

	printslice(names)

}

/*type vertex struct {
	X float64

	Y float64
}

func (v vertex) abs() float64 {

	return math.Sqrt(v.X*v.X + v.Y*v.Y)

}
func main() {

	v := vertex{
		5, 6,
	}

	fmt.Println(v.abs())

}
*/
