package main

import (
	"fmt"
	"strings"
)

func main() {
	// data := "apple,banana,orange,mango"
	// parts := strings.Split(data, ",")
	// fmt.Println(parts)

	// str := "one two two two one two three three three one two"
	// count := strings.Count(str, "two")
	// fmt.Println(count)

	str1 := "hello"
	str2 := "world"

	result := strings.Join([]string{str1, "Beautiful",  str2} , " ")
	fmt.Println(result)
}