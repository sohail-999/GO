package main

import "errors"

func divide(a, b float64) (result float64, err error) {
	if b == 0 {
		err = errors.New("division by zero")
		return // bare return — sends current values of result and err
	}
	result = a / b
	return // bare return
}

func main() {
	res, err := divide(10, 2)
	if err != nil {
		println("Error:", err.Error())
		return
	}
	println("Result:", res)

	res, err = divide(10, 0)
	if err != nil {
		println("Error:", err.Error())
		return
	}
}
