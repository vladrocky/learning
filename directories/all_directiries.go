package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	//PrintAllFiles("C:/WORKSPACE/PETS/learning/maps")
	PrintAllFilesWithFilter("C:/WORKSPACE/PETS/learning/maps", "task")
}

func PrintAllFiles(path string) {
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
		// печатаем имя элемента
		fmt.Println(filename)
		// если элемент — директория, то вызываем для него рекурсивно ту же функцию
		if f.IsDir() {
			PrintAllFiles(filename)
		}
	}
}

func PrintAllFilesWithFilter(path string, filter string) {
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
		if filename := filepath.Join(path, f.Name()); strings.Contains(filename, filter) {
			// печатаем имя элемента
			fmt.Println(filename)
			// если элемент — директория, то вызываем для него рекурсивно ту же функцию
			if f.IsDir() {
				PrintAllFiles(filename)
			}
		}

	}
}
