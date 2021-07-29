package main

import (
	"fmt"
	"github.com/ThompsonJonM/go-learning/testing/saying"
)

func main() {
	s := saying.Greet("James")
	fmt.Println(s)
}
