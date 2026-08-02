// @write a program to reads the marks of the student and then prints its grade A/B/C/D/FAIL
// first get the marks of the student
// aplly the condition to check the grade
package main

import "fmt"

func main() {

	var studentgarde uint
	fmt.Println("Enter the student marks:")
	fmt.Scan(&studentgarde)
	if studentgarde >= 85 {
		fmt.Println("The student achieves the A Grade")
	} else if studentgarde <= 85 && studentgarde >= 75 {
		fmt.Println("The student achieves the B Grade")
	} else if studentgarde <= 75 && studentgarde >= 65 {
		fmt.Println("The student achieves the C Grade")
	} else if studentgarde <= 65 && studentgarde >= 50 {
		fmt.Println("The student achieves the D Grade")
	} else {
		fmt.Println("Unfortunately , the student failed")
	}

}
