package main

import "fmt"

func main() {

	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else if x > 5 && x <= 10 {
		fmt.Println("x is between 5 and 10")
	} else {
		fmt.Println("x is less than 5")
	}
}