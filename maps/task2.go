package main

// Ассоциативные массивы широко применяются при решении алгоритмических задач.
// Когда количество данных более нескольких десятков, поиск значения в map происходит эффективнее, чем в массиве.
// Опираясь на эту информацию, попробуйте решить следующую задачу, которую часто предлагают на собеседованиях.
// Дан массив целых чисел A и целое число k. Нужно найти и вывести индексы пары чисел, сумма которых равна k.
// Если таких чисел нет, то вернуть пустой слайс. Индексы можно вернуть в любом порядке.

import "fmt"

func main() {

	A := []int{2, 1, 8, 5, 3, 4, 5}
	k := 10

	// mapA := make(map[int]int)

	// for _, val := range A {
	// 	mapA[val] = k - val
	// }

	// for i := 0; i < len(A); i++ {
	// 	for j := i + 1; j < len(A); j++ {
	// 		if A[i]+A[j] == k {
	// 			fmt.Println(i, j)
	// 		}
	// 	}
	// }

	result := find(A, k)

	fmt.Println(result)
}

func find(inp []int, k int) []int {

	mapA := make(map[int]int)

	for key, val := range inp {

		if keyMap, op := mapA[k-val]; op {
			//fmt.Println(keyMap)
			//fmt.Println(op)
			return []int{key, keyMap}
		}

		mapA[val] = key

	}

	return nil
}
