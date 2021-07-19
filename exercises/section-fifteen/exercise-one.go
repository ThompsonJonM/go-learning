package main

import "fmt"

func main() {
	x := foo()
	y, z := bar()

	fmt.Println(x, y, z)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 1984, "Hello world"
}
