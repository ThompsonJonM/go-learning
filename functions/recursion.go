package main

import "fmt"

func main() {
	// Anything done with recursion can be done with a loop
	n := factorial(4)

	fmt.Println(n)

	n1 := factorialForLoop(4)
	fmt.Println(n1)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factorialForLoop(n int) int {
	total := 1

	for ; n > 0; n-- {
		total *= n
	}

	return total
}
