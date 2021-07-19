package main

import "fmt"

// Package-level scope
var x int

func main() {
	x = 42
	fmt.Println(x)
	foo()

	a := incrementor()
	b := incrementor()
	fmt.Println(a()) // 1
	fmt.Println(a()) // 2
	fmt.Println(b()) // 1
}

func foo() {
	x++
	fmt.Println(x)
}

func incrementor() func() int {
	// Function-level scope
	var y int
	return func() int {
		y++
		return y
	}
}
