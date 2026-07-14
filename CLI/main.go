package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	command := flag.String("command", "", "command to execute")
	flag.Parse()
	if *command == " " {
		fmt.Println("Please provide a command")
		flag.Usage()
		return

	}
	switch *command {

	case "greet":
		fmt.Println("hello bhaijaan")
	case "time":
		fmt.Println("Current time", time.Now().Format("15:04:05"))
	default:
		fmt.Println("Unknown Command:", command)
	}

}
