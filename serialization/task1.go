package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name        string
	Email       string
	DateOfBirth time.Time
}

func main() {
	person := Person{
		Name:  "test",
		Email: "test@email",
	}

	str, err := serializToJSON(&person)
	fmt.Println(str, err)
}

func serializToJSON(input *Person) (string, error) {

	var out []byte
	err := json.Unmarshal(out, &input)

	fmt.Println(out)
	return string(out), err
}
