package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	p := Person{
		Name: "Ali",
		Age:  22,
		City: "faislabad",
	}

	jsonData, _ := json.MarshalIndent(p, "", "  ")
	os.WriteFile("data.json", jsonData, 0644)

}
