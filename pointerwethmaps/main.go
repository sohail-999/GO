package main

import (
	"fmt"
)

func maprip(m *map[string]int) { //m is the pointer to a map

	*m = map[string]int{
		"fakevalue": 100}
}

func main() {

	scores := map[string]int{
		"realkey": 10}

	maprip(&scores)

	fmt.Println(scores)

}
