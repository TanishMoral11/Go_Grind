package main

import "fmt"

func main() {
// var num int
// num = 10


// var ptr *int
// ptr = &num

num := 2
ptr := &num

fmt.Println("Value of num is", num)
fmt.Println("Address of num is", ptr)
fmt.Println("Data contains through pointer is", *ptr)
}