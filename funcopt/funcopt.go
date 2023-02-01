package main

import (
	"fmt"
	"math"
)

type figures int

const (
	square   figures = iota // квадрат
	circle                  // круг
	triangle                // равносторонний треугольник
)

func main() {
	ar, ok := area(2)
	if !ok {
		fmt.Println("Ошибка")
		return
	}
	myArea := ar(2)

	fmt.Println(myArea)
}

func area(fig figures) (func(float64) float64, bool) {

	var function func(x float64) float64
	var isH bool

	switch {

	case fig == 1:
		function = func(x float64) float64 {
			return x * x
		}
		isH = true
	case fig == 2:
		function = func(radius float64) float64 {
			return 2 * math.Pi * radius * radius
		}
		isH = true
	case fig == 2:
		function = func(x float64) float64 {
			return math.Sqrt(3) / 4 * x * x
		}
		isH = true
	case fig <= 0 && fig >= 3:
		function = nil
	}
	return function, isH
}
