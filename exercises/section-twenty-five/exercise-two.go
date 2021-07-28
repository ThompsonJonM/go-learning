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

	b, err := toJson(p1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))
}

func toJson(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return nil, fmt.Errorf("a json error occurred: %v", err)
	}
	return bs, nil
}
