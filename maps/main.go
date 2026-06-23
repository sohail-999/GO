package main

import "fmt"

func main() {

	fruitsprices := map[string]int{ //string is key type and int is the value type

		"Apple":      50,
		"Banana":     20,
		"Grapes":     90,
		"Watermelon": 100,
	}

	fmt.Println(fruitsprices)
}

//and map is works like table which is hash table
//and the does not follow the order and works on values and keys and which makes it faster than array in some operations
