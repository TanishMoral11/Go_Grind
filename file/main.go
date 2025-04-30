package main

import (
	"fmt"
	"os"
)

func main() {

	// file , err := os.Create("example.txt")

	// if err != nil {
	// 	fmt.Println("Error While Creating File:", err)
	// 	return
	// }
	// defer file.Close()

	// content := "Hello World By Tanish Moral"
	// _  ,errors := io.WriteString(file, content+"\n")
	// if errors != nil {
	// 	fmt.Println("Error While Writing to File:", errors)
	// 	return
	// }

	// fmt.Println("File Created and Content Written Successfully")

	/*
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error While Opening File:", err)
		return
	}
	defer file.Close()

	// create a buffer to read the file content
	buffer := make([]byte  , 1024)

	// read the file content into the buffer

	for {
		n , err := file.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error While Reading File:", err)
			return
		}

		fmt.Println(string(buffer[:n]))	}

		*/

		// read the entire file into a byte slice

		content ,err := os.ReadFile("example.txt")

		if err != nil {
			fmt.Println("Error While Reading File: ", err)
			return
		}

		fmt.Println(string(content))
}