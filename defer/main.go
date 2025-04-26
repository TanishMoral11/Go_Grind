package main

import "fmt"

func add(x, y int) int {
	return x+y
}

func main() {
	result := add(7,7)
	defer fmt.Println("The result is:", result)
	fmt.Println("Starting the program")
	defer fmt.Println("Middle of the program")
	fmt.Println("Ending the program")
}