package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)
	f1 := bar()
	fmt.Println(f1())
	fmt.Println(bar()()) // Execute bar, then execute anonymous function returned from bar
}

func foo() string {
	return "Hello world"
}

func bar() func() int {
	return func() int {
		return 451
	}
}
