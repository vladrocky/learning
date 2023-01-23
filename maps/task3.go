package main

import (
	"fmt"
)

func main() {
	// Представьте, что на входе есть слайс строк:
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}

	result := RemoveDuplicates(input)
	fmt.Println(result)
}

// Напишите функцию, которая убирает дубликаты, сохраняя порядок слайса:
func RemoveDuplicates(input []string) []string {

	mapA := make(map[string]bool)
	arr := make([]string, 0)

	for _, val := range input {

		if _, op := mapA[val]; !op {
			arr = append(arr, val)
			mapA[val] = true
		}
	}

	return arr
}
