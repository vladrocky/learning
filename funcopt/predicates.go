package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// containsDot возвращает все пути, содержащие точки
func containsDot(s string) bool {
	return strings.Contains(s, "2")
}

func PrintAllFilesWithFuncFilter(path string, predicate func(string) bool) {
	// создаём переменную, содержащую функцию обхода
	// мы создаём её заранее, а не через оператор :=, чтобы замыкание могло сослаться на него
	var walk func(string)
	walk = func(path string) {
		// получаем список всех элементов в папке (и файлов, и директорий)
		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("unable to get list of files", err)
			return
		}
		//  проходим по списку
		for _, f := range files {
			// получаем имя элемента
			// filepath.Join — функция, которая собирает путь к элементу с разделителями
			filename := filepath.Join(path, f.Name())
			// печатаем имя элемента, если путь к нему содержит filter, который получим из внешнего контекста
			if predicate(filename) {
				fmt.Println(filename)
			}
			// если элемент — директория, то вызываем для него рекурсивно ту же функцию
			if f.IsDir() {
				walk(filename)
			}
		}
	}
	// теперь вызовем функцию walk
	walk(path)
}

func main() {
	PrintAllFilesWithFuncFilter("C:/WORKSPACE/PETS/learning/maps", containsDot)
}
