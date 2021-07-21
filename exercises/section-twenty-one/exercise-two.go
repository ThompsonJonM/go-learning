package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello, my name is", p.first, p.last)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Jonathan",
		last:  "Thompson",
		age:   34,
	}

	p1.speak()
	// saySomething(p1) Compiler error: Requires a pointer
	saySomething(&p1) // No compiler error
}
