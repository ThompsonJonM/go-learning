package main

import "fmt"

func main() {
	defer foo() // Will run when containing function exits
	bar()
}

func foo() {
	fmt.Println("Hello world")
}

func bar() {
	fmt.Println("Bar")
}
