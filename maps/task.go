package main

import "fmt"

func main() {
	price_map := map[string]int{
		"хлеб":     50,
		"молоко":   100,
		"масло":    200,
		"колбаса":  500,
		"соль":     20,
		"огурцы":   200,
		"сыр":      600,
		"ветчина":  700,
		"буженина": 900,
		"помидоры": 250,
		"рыба":     300,
		"хамон":    1500,
	}

	for key, val := range price_map {
		if val >= 500 {
			fmt.Println(key)
		}
	}

	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	order_cost := 0

	for _, item := range order {
		order_cost += price_map[item]
	}

	fmt.Println(order_cost)
}
