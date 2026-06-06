package main

import "fmt"

func main() {

	fruitsprice := map[string]int{ //string is key type and int is the value type

		"apple":      50,
		"banana":     20,
		"grape":      90,
		"watermelon": 100,
	}

	fmt.Println(fruitsprice)
}

//and map is works like table which is hash table
//and the does not follow the order and works on values and keys and which makes it faster than array in some operations
