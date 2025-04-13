package main

import "fmt"

func main() {

	fmt.Println("I'm Learning Array in Golang")

	var name[5] string
	name[0] = "John"
	name[2] = "Doe"

	fmt.Println("Names of the persons are : ", name)

	var numbers = [5]int{1,2,3,4,5}
	fmt.Println("Numbers are : ", numbers)
	fmt.Println("Length of the array is : ", len(numbers))

	var price[5] string
	fmt.Print("Prices are : ", price)
}