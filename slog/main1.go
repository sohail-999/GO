package main

import (
	_ "embed"
	"fmt"
)

var (
	//go:embed test.txt
	Text_File []byte
)

func main() {

	fmt.Println(string(Text_File))

}
