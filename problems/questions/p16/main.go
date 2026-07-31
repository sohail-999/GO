// @write a program to read a charactor  wheather it is alphabet,digit or the special charactotrs
// first get the alphabet/digit/special symbol from the user
// apply condition to check
package main

import "fmt"

func main() {
	var character rune
	fmt.Println("Please enter a  charactor:")
	fmt.Scanf("%c", &character)
	if (character >= 'A' && character <= 'Z') || (character >= 'a' && character <= 'z') {
		fmt.Printf("%c is a Alphabet\n", character)
	}
	if character >= '0' && character <= '9' {

		fmt.Printf("%c is a Digit\n", character)
	} else {

		fmt.Printf("%c is a special symbol\n", character)

	}

}
