package main

import (
	"fmt"
	"github.com/ThompsonJonM/go-learning/dog"
)

type canine struct {
	name string
	age  float64
}

func main() {
	d := canine{
		name: "Winston",
		age:  dog.Years(0.9),
	}

	fmt.Println(d)
}
