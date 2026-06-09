package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {

	v := Vertex{3, 2}
	p := &v
	p.X = 1e2 //1e10 shows that 1 mutiply by 10 to the power of 9
	fmt.Println(v)

}
