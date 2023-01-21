package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Удаление последнего элемента слайса
	s := []int{1, 2, 3}
	if len(s) != 0 { // защищаемся от паники
		s = s[:len(s)-1]
	}
	fmt.Println(s) // [1 2]

	// Удаление первого элемента слайса:
	s1 := []int{1, 2, 3}
	if len(s1) != 0 { // защищаемся от паники
		s1 = s1[1:]
	}
	fmt.Println(s1) // [2 3]

	// Удаление элемента слайса с индексом i:
	s2 := []int{1, 2, 3, 4, 5}
	i := 2

	if len(s2) != 0 && i < len(s2) { // защищаемся от паники
		s2 = append(s2[:i], s2[i+1:]...)
	}
	fmt.Println(s2) // [1 2 4 5]

	// Сравнение двух слайсов:
	s11 := []int{1, 2, 3}
	s22 := []int{1, 2, 4}
	s3 := []string{"1", "2", "3"}
	s4 := []int{1, 2, 3}

	fmt.Println(reflect.DeepEqual(s11, s22)) // false
	fmt.Println(reflect.DeepEqual(s11, s3))  // false
	fmt.Println(reflect.DeepEqual(s11, s4))  // true
}
