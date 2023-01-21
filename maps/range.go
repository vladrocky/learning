package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["foo"] = "bar"
	m["bazz"] = "yup"
	for k, v := range m {
		// k будет перебирать ключи,
		// v — соответствующие этим ключам значения
		fmt.Printf("Ключ %v, имеет значение %v \n", k, v)
	}

	// предложение
	sentence := "Πολύ λίγα πράγματα συμβαίνουν στο σωστό χρόνο, και τα υπόλοιπα δεν συμβαίνουν καθόλου"
	// инициализируем map
	// ключами будут знаки, а значениями — количество вхождений
	frequency := make(map[rune]int)
	// пройдёмся по знакам в предложении
	for _, v := range sentence {
		frequency[v]++ // встреченному знаку увеличиваем частоту
	}
	// печатаем
	for k, v := range frequency {
		fmt.Printf("Знак %c встречается %d раз \n", k, v)
	}
}
