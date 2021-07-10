package main

import "fmt"

var y string

func main() {
	fmt.Printf("%T\n", y)
	fmt.Println("Printing the value of y:", y, "ending.")

	y = "This is a string"

	fmt.Printf("%T\n", y)
	fmt.Println("Printing the value of y:", y, "ending.")
}
