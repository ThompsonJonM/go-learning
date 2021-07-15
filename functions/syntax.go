package main

import "fmt"

func main() {
	foo()
	bar("Sarah")

	s1 := zed("James")
	fmt.Println(s1)

	x, y := newFunc("Jonathan", "Thompson")
	fmt.Println(x, y)
}

// func (r receiver) identifier(parameters) (return(s)) { ... }
func foo() {
	fmt.Println("Hello world")
}

// Everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello,", s)
}

func zed(s string) string {
	return fmt.Sprint("Hello from zed, ", s)
}

func newFunc(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := false

	return a, b
}
