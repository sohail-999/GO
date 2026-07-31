//@write a program to read the charactors whethor its vowel or the consonent.
//first read/get the charactors from the user
//apply the conditions to check the vowel and the consonent

package main

import "fmt"

func main() {
	var vowel rune

	fmt.Println("Please enter a character:")
	fmt.Scanf("%c", &vowel)

	if vowel == 'a' || vowel == 'e' || vowel == 'i' || vowel == 'o' || vowel == 'u' ||
		vowel == 'A' || vowel == 'E' || vowel == 'I' || vowel == 'O' || vowel == 'U' {
		fmt.Printf("%c is a vowel\n", vowel)
	} else {
		fmt.Printf("%c is not a vowel\n", vowel)
	}
}
