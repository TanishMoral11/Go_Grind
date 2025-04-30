package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Learning web Services in Golang")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error While Making Request:", err)
		return
	}
	fmt.Println("Type of Response:", res)
}