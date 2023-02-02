package main

import "fmt"

var Global int = 5

func main() {

	defer func(a int) {
		fmt.Printf("\ndefer Global varible = %d", a)
	}(Global)

	Global++
	fmt.Printf("Global varible = %d", Global)
}
