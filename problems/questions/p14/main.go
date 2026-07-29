// Write a program that  reads a year and check whether  its leap year or not
// first get the year as input from the user
// then apply the some condition
// leap year is divsible by 4 or divisible by 4,100 and 400 otherwise its not leap year
package main

import "fmt"

func main() {

	var year int
	fmt.Println("Please enter the year:")
	fmt.Scan(&year)
	leapyear := year

	if leapyear%4 == 0 && leapyear%100 != 0 {
		fmt.Printf("%v is a leap year\n", leapyear)
		//fmt.Printf("%v is a leap year\n", year)

	}
	if leapyear%4 != 0 {
		fmt.Printf("%v is not a leap year\n", leapyear)
	}
	if leapyear%4 == 0 && leapyear%100 == 0 && leapyear%400 == 0 {
		leapyear = year
		fmt.Printf("YES , %v is a leapyear\n", leapyear)
		//fmt.Printf("%v is leap year\n", year)
	}
	if leapyear%4 == 0 && leapyear%100 == 0 && leapyear%400 != 0 {
		leapyear = year
		fmt.Printf("NO , %v is not a leapyear\n", leapyear)
		//fmt.Printf("%v is leap year\n", year)
	}

}
