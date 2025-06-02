package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	IsAdult bool `json:"is_adult"`
}
func main() {
	fmt.Println("We Are Learning JSON in Go")

	person := Person{
		Name: "John Doe",
		Age: 30,
		IsAdult: true,
	}

	// fmt.Println("Person Data : ", person)

	//Convert person into JSON Encoding (Marshalling)

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error Marshalling JSON: ", err)
		return
	}

	fmt.Println("Json Data : ", string(jsonData))

	// Decoding (Unmarshalling)

	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error Unmarshalling JSON: ", err)
		return
	}

	fmt.Println("Unmarshalled Data : ", personData)
}