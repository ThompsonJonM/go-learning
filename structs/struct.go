package main

import "fmt"

// Data structure allowing composition of values of different types
// Also known as aggregate or complex data structure
type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(p1, p2)
	fmt.Println(p1.first, p1.last) // Dot notation for calling specific values
}
