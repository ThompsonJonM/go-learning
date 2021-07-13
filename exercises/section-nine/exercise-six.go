package main

import "fmt"

func main() {
	x := "James Bond"
	if x == "Moneypenny" {
		fmt.Println("Good to see you, Moneypenny")
	} else if x == "M" {
		fmt.Println("Hello, M")
	} else {
		fmt.Println("Bond, James Bond")
	}
}
