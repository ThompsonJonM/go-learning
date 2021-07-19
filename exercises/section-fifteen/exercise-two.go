package main

import "fmt"

func main() {
	x := []int{2, 3, 4}

	n1 := foo(x...)
	n2 := bar(x)
	fmt.Println(n1, n2)
}

func foo(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}

	return sum
}

func bar(n []int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}

	return sum
}
