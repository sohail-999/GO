package main

import (
	"fmt"
)

func sendticket(usertickets uint, firstname string, lastname string, email string) {

	var ticket = fmt.Sprintf("%v sending ticket: %v %v", usertickets, firstname, lastname)
	fmt.Println("##############################")
	fmt.Printf("Sending Ticket \n %v \n to the email address %v\n", ticket, email)
	fmt.Println("##############################")
}

func main() {

	var Usertickets uint = 50
	Firstname := "John"
	Lastname := "Doe"
	Email := "john.doe@example.com"

	fmt.Println("Welcome to the Go Conference booking application\nWe have a total of", Usertickets, "tickets available for the Person Name", Firstname, Lastname, "with Email User provided", Email)

	sendticket(Usertickets, Firstname, Lastname, Email)

}
