package main

import "fmt"

func double(nums []int) {

	for i := range nums {

		nums[i] *= 2 // nums[i] = nums[i] * 2.

	}
}

func main() {

	s := []int{1, 2, 4}

	double(s)

	fmt.Println(s) // [2 4 6]
}
