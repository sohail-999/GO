//Write a program that  reads a year and check whether  its leap year or not
//first get the year as input from the user
//then apply the some condition

package main

import "fmt"

func main() {

	var year int
	fmt.Println("Please enter the number:")
	fmt.Scan(&year)
	leapyear := year

	if leapyear%4 == 0 && leapyear%100 != 0 {
		fmt.Println("this is a leap year")
		//fmt.Printf("%v is a leap year\n", year)

	}
	if leapyear%4 == 0 && leapyear%100 == 0 && leapyear%400 == 0 {
		leapyear = year
		fmt.Printf("YES , %v is a leapyear\n", leapyear)
		//fmt.Printf("%v is leap year\n", year)
	}

}
