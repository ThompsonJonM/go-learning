package main

import "fmt"

func main() {
	var x [5]int
	var y [6]int // NOT the same type as x as len(y) is part of type

	fmt.Println(x)
	x[3] = 42

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(len(x))
}
