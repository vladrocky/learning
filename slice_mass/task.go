package main

import "fmt"

func main() {

	slice := make([]int, 0, 100)

	for index := range slice {
		slice[index] = index + 1
	}

	new_slice := append(slice[:10], slice[90:]...)

	sl_len := len(new_slice)

	new_slice2 := make([]int, sl_len)

	for index := range new_slice2 {
		new_slice2[index] = new_slice[sl_len-index-1]
	}

	fmt.Println(new_slice2)
}
