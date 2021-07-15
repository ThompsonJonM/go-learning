package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person  // embedded type, promoted to the outer type; also known as anonymous field
	license bool
}

func main() {
	sa1 := secretAgent{
		person: person{ // unqualified type name ACTS as the field name
			first: "James",
			last:  "Bond",
			age:   32,
		},
		license: true,
	}

	fmt.Println(sa1)
	fmt.Println(sa1.first)        // Promoted to the top level of the type
	fmt.Println(sa1.person.first) // Could also be used in case of name collisions
}
