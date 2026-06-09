package main

import "fmt"

func replaceMap(m *map[string]int) {

	*m = map[string]int{
		" x": 99}
	// replaces the entire map
}

func main() {

	scores := map[string]int{"Ali": 10}

	replaceMap(&scores)

	fmt.Println(scores)

	// map[x:999]  ← whole map swapped!
}
