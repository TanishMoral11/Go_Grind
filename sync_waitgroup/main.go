package main

import (
	"fmt"
	"sync"
)

func worker(i int ,wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("Worker %d started\n", i)
	// Some Task is happening
	fmt.Printf("Worker %d Ended\n", i)
}

func main() {

	fmt.Println("Explore goroutine started")

	var wg sync.WaitGroup

	// Start 3 worker goroutines
	for i:= 1; i<=3; i++{
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("Explore goroutine ended")
}

