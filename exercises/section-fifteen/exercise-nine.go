package main

import "fmt"

func main() {
	s := bar(foo)

	fmt.Println(s)
}

func foo() string {
	return "Hello world"
}

func bar(f func() string) string {
	return f()
}
