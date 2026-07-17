//write a program to read the radius of the circle and then prints the area and circumference
//give the raduis number
//apply the circumference formula and print it
//apply the  area formula then prints it

package main

import "fmt"

func circumferenceandarea(radius float64) {
	fmt.Println("Circumference:", 2*3.14*radius) //formula (2*pi*radius)
	fmt.Println("Area:", 3.14*radius*radius)     //formula(pi*r^2)

}

func main() {
	circumferenceandarea(7)
}
