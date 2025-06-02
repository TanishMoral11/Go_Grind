package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func performGetRequest() {
	// res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	// if err != nil {
	// 	fmt.Println("Error fetching data:", err)
	// 	return
	// }

	// defer res.Body.Close()

	// if res.StatusCode != http.StatusOK {
	// 	fmt.Printf("Error: received status code %d\n", res.StatusCode)
	// }

	// // data, err := ioutil.ReadAll(res.Body)
	// // if err != nil {
	// // 	fmt.Println("Error reading response body:", err)
	// // 	return
	// // }
	// // fmt.Println("Data: ", string(data))

	// var todo Todo
	// // Decoder : it decodes into normal objects and saves into in this todo variable 
	// err = json.NewDecoder(res.Body).Decode(&todo)
	// if err != nil {
	// 	fmt.Println("Error decoding JSON:", err)
	// 	return
	// }
	// fmt.Println("Todo Item: ", todo)

	// fmt.Println("Title : ", todo.Title)
	// fmt.Println("Completed : ", todo.Completed)
}

func performPostRequest() {
	

	todo := Todo{
		UserID: 69, 
		Title: "Tanish Moral",
		Completed: true,
	}

	// Convert The Todo struct to JSON
	jsonData , err := json.Marshal(todo)
	if err != nil {	
		fmt.Println("Error converting to JSON:", err)
		return
	}

	// Convert the JSON data to a string
	jsonString := string(jsonData)

	// Convert the String data to a reader
	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos"
	res , err  := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}


	defer res.Body.Close()

	data , _ := ioutil.ReadAll(res.Body)
	fmt.Println("Resonse ", string(data))
}


func main() {
	fmt.Println("Learning Go with CRUD operations\n")
	// performGetRequest()
	performPostRequest()

}
