package main

import (
	"fmt"
	"time"
)

func sayHello(){
	fmt.Println("Hello, Goroutine!")
	time.Sleep(2*time.Second)
	fmt.Println("Goodbye, Goroutine!")
}

func sayHi(){
	fmt.Println("Hi!!, Tanish Moral ")
}

func main() {
	fmt.Println("Learning Goroutine")
	go sayHello()
	sayHi()
}