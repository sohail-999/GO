//@write a program to read the seconds and convert them into hours , minutes ,seconds
//first get the seconds
//convert the seconds into minutes
//convert them jnto hours

package main

import "fmt"

//func convertor(a int) {
//	fmt.Println("Seconds:\n",a)
//}

func main() {
	var a uint
	fmt.Println("Write the Seconds:")
	fmt.Scan(&a)
	fmt.Println("Seconds into Minutes:\n", a/60, "minutes")
	fmt.Println("Seconds into hours:\n", a/3600, "hours")
}
