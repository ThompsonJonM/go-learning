package main

import "fmt"

var y = 42
var z string = "This is a string"
var a string = `this is a "raw" string`

type newString string

var b newString = "Hello World"

func main() {
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
	fmt.Println(z, "and", a)
	fmt.Printf("%T\n", b)

	a = string(b) // conversion, not casting
	fmt.Println(a)
	fmt.Printf("%T", a)

	// The following assignment will fail unlike in JS due to the compiler checking for
	// type safety. Var z is declared as a string and ONLY a string. It can never be
	// another type.
	// z = 43
	// fmt.Println(z)
}
