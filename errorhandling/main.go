package main

import (
	"fmt"
	"os"
)

func main() {

	// @ for opening the file
	file, err := os.Open("Daata.txt") //the os allows to interact with opearting system

	if err != nil {
		fmt.Println("Error :", err) //fmt.Println("Error:", err) //if the file does not exists its gonna print

		return
	}

	fmt.Println("File opened successfully:", file.Name()) //if the file exists in the folder it this gonna print

	//@ for reading the file

	// 	file,err :=os.ReadFile("Data.txt")

	// 	if err!=nil{

	//	      fmt.Println("Error:", err)
	//	      return
	//		}
	//
	// fmt.Println("File readed suucessfully:",string(file))
}
