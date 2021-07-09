package main

import "fmt"

// Package level scope
var y = 43
var z int

func main() {
	// Function level scope
	// Try to limit scope as often as possible by using short declaration
	// operator inside of functions rather than package level.
	x := 42
	fmt.Println(x, y)
	foo()
	fmt.Println(z) // should return zero value
}

func foo() {
	fmt.Println("Running again:", y)
}
