package main

import "fmt"

func main() {
	defer hello()
	n := bar()

	fmt.Println(n)
}

func hello() {
	fmt.Println("Hello world")
}

func bar() int {
	return 42
}
