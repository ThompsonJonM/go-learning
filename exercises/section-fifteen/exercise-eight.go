package main

import "fmt"

func main() {
	x := bar()
	y := x()

	// Three separate ways of calling
	fmt.Println(y)       // As a var
	fmt.Println(x())     // Call x
	fmt.Println(bar()()) // Call original func twice
}

func foo() string {
	return "Hello world"
}

func bar() func() string {
	return func() string {
		return "Hello world"
	}
}
