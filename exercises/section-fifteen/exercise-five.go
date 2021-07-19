package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{
		length: 2.0,
		width:  2.0,
	}

	c := circle{
		radius: 3.0,
	}

	a1 := s.area()
	a2 := c.area()
	fmt.Println(a1, a2)
}
