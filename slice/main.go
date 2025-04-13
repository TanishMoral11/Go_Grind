package main

import "fmt"

func main() {

	num := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice of numbers is : ", num)
	fmt.Printf("Type of num is %T\n", num)

	stock := make([]int , 10)
	fmt.Println("Slice of stock is : ", stock)
	fmt.Printf("Type of stock is %T\n", stock)
	fmt.Print("Length of stock is : ", len(stock), "\n")
	fmt.Println("Capacity of stock is : ", cap(stock), "\n")
}