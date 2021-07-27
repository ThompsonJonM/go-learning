package main

import "fmt"

func main() {
	// Always, always, ALMOST ALWAYS check your errors
	var answer1 string

	fmt.Print("Enter your name: ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer1)
}
