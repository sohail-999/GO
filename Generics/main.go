package main

import (
	"fmt"
	"math"
)

type vertex struct {
	X float64
	Y float64
}

func (v vertex) abs() float64 {

	return math.Sqrt(v.X*v.X + v.Y*v.Y)

}
func (v *vertex) scale(f float64) { ///The Scale method must have a pointer receiver to change the Vertex value declared in the main function.

	v.X = v.X * f
	v.Y = v.Y * f
}
func main() {

	v := vertex{
		4,
		4,
	}
	v.scale(10)
	fmt.Println(v.abs())

}
