package main

import "fmt"

const (
	a         = 42    // Constant "of a kind"
	b float64 = 42.78 // Typed constant
	c         = "Hello world"
)

func main() {
	fmt.Println(a, b, c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
