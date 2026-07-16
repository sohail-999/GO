package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "My New CLI Application" //name
	app.Action = (func(ctx *cli.Context) error {
		fmt.Println("app launched")
		return nil

	})
	err := app.Run(os.Args) //for accepting the arg from the terminal
	if err != nil {

		log.Fatal(err)

	}

}
