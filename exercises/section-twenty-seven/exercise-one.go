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
		age:  0.9,
	}

	y := dog.Years(d.age)
	fmt.Println(y)
}
