package main

import "fmt"

func main() {
	m := map[int]string{1: "first"}
	v, ok := m[1]
	fmt.Println(v, ok)
	delete(m, 2) // ошибки не будет
	delete(m, 1)
	v, ok = m[1]
	fmt.Println(v, ok)
}
