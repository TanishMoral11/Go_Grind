package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println(currentTime)

	formattedtime := currentTime.Format("02-01-2006, Monday, 15:04:05");
	fmt.Println("Formatted Time : ", formattedtime);
}