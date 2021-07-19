package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	p.first = "James"
	p.last = "Bond"
	p.age = 32
}

func main() {
	p := person{
		first: "Jonathan",
		last:  "Thompson",
		age:   34,
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}
