package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

// Non-Pointer receiver works with both pointer and non-pointer VALUES
// Pointer receiver ONLY works with pointer VALUES
func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("area:", s.area())
}

func main() {
	c := circle{
		radius: 5,
	}
	info(&c)
}
