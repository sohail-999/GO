package main

import "fmt"

func main() {

	var num int = 14

	fmt.Println("Here is your permission for a drive:")
	fmt.Println("System checking your age...")

	if num >= 18 {

		fmt.Println("you r permitted for a drive")

	} else if num >= 14 && num <= 17 {

		fmt.Println("you r  permitted for a drive BUT with your parents")

	} else {

		fmt.Println("you r not permitted for a drive")

	}
}
