package main

import "fmt"

func main() {

	const ( //the frormula is 2 power n.
		KB = 1 << (10 * (iota + 1)) // 1 << 10 = 1024
		MB                          // 1 << 20 = 1048576
		GB                          // 1 << 30 = 1073741824
	) // 6

	fmt.Println("Here is the size of KB, MB and GB in bytes:")

	fmt.Println(KB, MB, GB)
}
