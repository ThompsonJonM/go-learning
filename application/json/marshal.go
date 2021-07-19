package main

import (
	"encoding/json"
	"fmt"
)

// Must capitalize fields in order to marshal
type person struct {
	First string
	Last  string
	Age   int
}

type secretAgent struct {
	person
	License bool
}

func main() {
	p1 := secretAgent{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   32,
		},
		License: true,
	}

	p2 := secretAgent{
		person: person{
			First: "Alec",
			Last:  "Trevelyan",
			Age:   38,
		},
		License: true,
	}

	people := []secretAgent{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(bs)
	fmt.Println(string(bs))
}
