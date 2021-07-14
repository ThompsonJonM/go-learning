package main

import "fmt"

func main() {
	x := make([]int, 10, 10) // Create a slice of zero vals with len/cap of 10
	fmt.Println(x, len(x), cap(x))

	x[0] = 1
	x[9] = 10

	x = append(x, 11)
	fmt.Println(x, len(x), cap(x)) // Cap now 20 as it doubles when max is reached
}
