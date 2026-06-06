package main

import "fmt"

func main() {
	ages := map[string]int{

		"Ali": 25, "Sara": 30, "Usman": 22,
	}

	for key, value := range ages {

		fmt.Println(key, value)

	}

}
