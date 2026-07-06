package main

import (
	"encoding/json"
	"fmt"
)

type myInt struct {
	Intvalue int `json:"intValue"`
}

func main() {
	data := &myInt{Intvalue: 1234}

	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("marshal error:", err)
		return
	}

	fmt.Println(string(b))
}
