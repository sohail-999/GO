package main

import "fmt"

func summarise(nums ...int) {

	total := 0

	for _, n := range nums {

		total += n
	}
	fmt.Printf("numbers : %v\n", nums)
	fmt.Printf("count   : %d\n", len(nums))
	fmt.Printf("total   : %d\n", total)
}

func main() {

	summarise(1, 2, 3)

	summarise(1, 2, 3, 4)

}
