package main

import "fmt"

func main() {
	// Composite literal: x := type{values}
	x := []int{2, 3, 4, 5}

	fmt.Println(x) // [2 3 4 5]
}
