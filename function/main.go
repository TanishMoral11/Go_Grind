package main

import "fmt"

func simpleFunction(){
	fmt.Println("I'm a simple function")
}

func add(a,b int)(int) {
	return a+b
}
func multiply(a, b int)(result int){
	result = a*b
	return
}

func main() {
	fmt.Println("I'm learning Golang")
	simpleFunction()
	add := add(1,2)
	fmt.Println("1 + 2 = ", add)
	mul := multiply(2, 3)
	fmt.Println("Multiplication of 2 * 3 = ", mul)
}