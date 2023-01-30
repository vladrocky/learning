package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Дана структура
// Напишите код, который будет сериализовывать структуру в json-строку
type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

func main() {
	person := Person{
		Name:        "test",
		Email:       "test@email",
		DateOfBirth: time.Now(),
	}

	str, err := serializToJSON(&person)
	fmt.Println(str, err)
}

func serializToJSON(input *Person) (string, error) {

	out, err := json.Marshal(&input)
	return string(out), err
}
