package main

import "fmt"

var x bool

func main() {
	a := 7
	b := 42

	fmt.Println(x) // false
	x = true
	fmt.Println(x) // true

	fmt.Println(a == b) // There is NO triple equal operator in Go
}
