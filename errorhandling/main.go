package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("Data.txt")

	if err != nil {
		fmt.Println("Error:", err) //if the file does not exists its gonna print
		return
	}

	fmt.Println("File opened successfully:", file.Name()) //if the file exists in the folder it this gonna print
}
