package main

import "fmt"

func divide (a, b int) (float64, error){
	if b==0 {
		return 0, fmt.Errorf("division by zero")
	}
	return float64(a) / float64(b), nil
}

func main() {

	fmt.Println("started error handling in the functions")
	div, _ := divide(10, 1)
	fmt.Println("Division of 10 / 0 = ", div)
}