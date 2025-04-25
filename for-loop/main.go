package main

import "fmt"

func main() {

	ct := 0
	for {
		ct++;
		fmt.Println("Infinite loop", ct)

	}

	// for i := 0; i<=10; i++ {
	// 	fmt.Println("Number is ", i)
	// }
	// for i := 0; i<40; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i, "is even")
	// 	} else if i%3 == 0 {
	// 		fmt.Println(i, "is odd and divisible by 3")
	// 	} else {
	// 		fmt.Println(i, "is odd and not divisible by 3")
	// 	}
	// }
}