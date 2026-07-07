package main

import (
	"fmt"
	"regexp"
)

func main() {

	//match, _ := regexp.MatchString("b([a-z]+)ake", "bananashake")
	r, _ := regexp.Compile("([a-z]+)ch") //regexp struct
	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindStringSubmatchIndex("peach  punch"))
	fmt.Println(r.ReplaceAllString("a peach", "banana"))
}
