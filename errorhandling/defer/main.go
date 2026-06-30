package main

import (
	"fmt"
	"io"
	"os"
)

func Copyfile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return

	}
	defer src.Close() //closes the file
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}

	defer dst.Close() //closes the file
	return io.Copy(dst, src)

}
func main() {

	written, err := Copyfile("source.txt", "destination.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Bytes written:", written) //prints the bytes of the file
}
