package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	license bool
}

// func (r receiver) identifier(parameter(s)) (return(s)) { ... }
// All instances of secretAgent now have access to the speak method
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		license: true,
	}

	fmt.Println(sa1)
	sa1.speak()
}
