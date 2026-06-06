package main

import "fmt"


func main(){

f := new(myFile)



f.fd = syscall.Open(...)

runtime.AddCleanup(f, func(fd int) {


	syscall.Close(f.fd) 
	// Mistake: We reference f, so this cleanup won't run!
}, f.fd)

}