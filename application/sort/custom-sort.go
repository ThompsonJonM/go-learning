package main

import (
	"fmt"
	"sort"
)

type person struct {
	First string
	Age   int
}

type ByName []person

// Apply methods necessary for sort Interface
func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].First < a[j].First }

func main() {
	p1 := person{
		First: "James",
		Age:   32,
	}

	p2 := person{
		First: "Jonathan",
		Age:   34,
	}

	p3 := person{
		First: "Timothy",
		Age:   36,
	}

	p4 := person{
		First: "Sarah",
		Age:   38,
	}

	people := []person{p4, p2, p1, p3}
	fmt.Println(people) // [{Sarah 38} {Jonathan 34} {James 32} {Timothy 36}]

	// Use conversion to apply ByName type to people struct,
	// then use Sort function against the ByName type
	sort.Sort(ByName(people))
	fmt.Println(people) // [{James 32} {Jonathan 34} {Sarah 38} {Timothy 36}]
}
