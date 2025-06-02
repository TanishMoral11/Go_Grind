package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func main() {
	fmt.Println("Learning Go with CRUD operations\n")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", res.StatusCode)
	}

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response body:", err)
	// 	return
	// }
	// fmt.Println("Data: ", string(data))

	var todo Todo
	// Decoder : it decodes into normal objects and saves into in this todo variable 
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println("Todo Item: ", todo)

	fmt.Println("Title : ", todo.Title)
	fmt.Println("Completed : ", todo.Completed)

}
