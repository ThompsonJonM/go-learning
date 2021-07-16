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

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last,"- the person speak")
}

// Interfaces allow us to define behavior and do polymorphism
type human interface {
	speak() // Any type that has method speak is ALSO of type human; a VALUE can be of more than one TYPE
}

func bar(h human) {
	switch h.(type) { // Switch on data TYPE
	case person:
		fmt.Println("I am a person,", h.(person).first, h.(person).last) // Example of TYPE assertion
	case secretAgent:
		fmt.Println("I am a secret agent,", h.(secretAgent).first, h.(secretAgent).last)
	}
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		license: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "No",
	}

	fmt.Println(sa1)
	sa1.speak()
	fmt.Println(p1)
	p1.speak()
	bar(sa1)
	bar(p1)
}
