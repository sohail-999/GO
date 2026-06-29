package main

import (
	"fmt"
)

func main() {

	//err := fmt.Errorf("the error occured %v :", time.Now())
	//fmt.Println("the error happened at the :", err)

	const name, depart = "geekforgeeks", "CS"
	err := fmt.Errorf("%q is a %q portal", name, depart)

	fmt.Println(err.Error())

}
