package main

import (
	"fmt"
	"regexp"
)

func main() {

	//match, _ := regexp.MatchString("b([a-z]+)ake", "bananashake")
	r, _ := regexp.Compile("b([a-z]+)akee")
	fmt.Println(r.MatchString("bananashakee"))

}
