package main

import "fmt"

func main() {
	// fmt.Printf("Enter A Number : ")
	// var num int
	// fmt.Scan(&num)
	// fmt.Println("You Entered ", num , "!!")

	fmt.Println("What's Your Name?")
	var name string
	fmt.Scan(&name)
	fmt.Println("Hello", name);
}