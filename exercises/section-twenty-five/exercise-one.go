package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	b, err := json.Marshal(p1)
	if err != nil {
		log.Fatalf("json error occurred: %v", p1)
	}
	fmt.Println(string(b))
}
