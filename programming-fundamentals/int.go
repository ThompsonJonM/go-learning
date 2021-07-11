package main

import "fmt"

var x int
var y float64
var z int8 = -128 // -129 causes compile error

func main() {
	x = 42
	// x = 2.354 Causes compile error due to type mismatch
	y = 42.34534

	fmt.Println(x, y, z)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
}
