package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learnig Url")
	myURL := "https://www.google.com/search?q=golang+url+package&rlz=1C1CHBF_enIN1010IN1010&oq=golang+url+package&aqs=chrome..69i57j0i512l9.10300j0j7&sourceid=chrome&ie=UTF-8"

	// fmt.Println("My Url is : ", myURL)
	fmt.Printf("myURL Type is %T\n", myURL)

	parsedURL , err := url.Parse(myURL)

	if err != nil {
		fmt.Println ("Error While Parsing URL:", err)
		return
	}

	fmt.Printf("Type of Parsed URL: %T\n", parsedURL)
	fmt.Println("Scheme: ", parsedURL.Scheme)
	fmt.Println("Host: ", parsedURL.Host)
	fmt.Println("Path: ", parsedURL.Path)
	fmt.Println("Raw Query: ", parsedURL.RawQuery)
}