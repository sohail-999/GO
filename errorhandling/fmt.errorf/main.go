package main

import (
	"errors"
	"fmt"
)

func main() {

	//@nested error handling for printing the time of the error

	//err := fmt.Errorf("the error occured %v :", time.Now())
	//fmt.Println("the error happened at the :", err)

	//@ printing the error using %q(which is used for putting the quotes)

	//const name, depart = "geekforgeeks", "CS"

	//err := fmt.Errorf("%q is a %q portal", name, depart)

	//fmt.Println(err.Error())
	originalerror := errors.New("original error")
	wrappederror := fmt.Errorf("additional information %w", originalerror)
	fmt.Println(wrappederror)
	//if errors.Is(wrappederror, originalerror) {
	//	fmt.Println("the wrapped contain the original...")

	// }
	unwrappederror := errors.Unwrap(wrappederror)
	fmt.Println(unwrappederror)

}
