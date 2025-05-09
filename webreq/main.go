package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	fmt.Println("Learning web Services in Golang")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error While Making Request:", err)
		return
	}

	defer res.Body.Close()
	fmt.Printf("Type of Response: %T\n", res)

	// Read the response body

	data, err := ioutil.ReadAll(res.Body)

	if err!= nil {
		fmt.Println("Error While Reading Response Body:", err)
		return
	}

	fmt.Println("Response : ", string(data));
}