//@write a program to read the charactors whethor its vowel or the consonent.
//first read/get the charactors from the user
//apply the conditions to check the vowel and the consonent

package main

import "fmt"

func main() {
	//var v1, v2, v3, v4, v5 rune
	var v rune
	fmt.Println("Please enter a charactor")
	fmt.Scan(&v)
	if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' ||
		v == 'A' || v == 'E' || v == 'I' || v == 'O' || v == 'U' {
		fmt.Printf("%v is a vowel /n", v)
	} else {
		fmt.Printf("%v is not a vowel \n", v)
	}
	//	vowels := 'a', 'e', 'i', 'o', 'u'

	/*v2 = 'e'
	v3 = 'i'
	v4 = 'o'
	v5 = 'u'*/
	//if v==a,e,i

}
