package main

import (
	"fmt"
	"github.com/ThompsonJonM/go-learning/exercises/section-twenty-nine/exercise-one/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	w := canine{
		name: "Winston",
		age:  1,
	}

	fmt.Println(w)
	fmt.Println(dog.Years(w.age))
}
