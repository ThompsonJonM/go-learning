package main

import "fmt"

func main() {
	i := []int{2, 3, 4, 5, 6, 7, 8}
	s := sum(i...)
	fmt.Println(s)

	s1 := even(sum, i...)
	s2 := odd(sum, i...)
	fmt.Println("Even numbers:", s1)
	fmt.Println("Odd numbers:", s2)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v

	}
	return total
}

// Example of a callback
// Passing a function, its parameters, and return type into another function
func even(f func(x ...int) int, v ...int) int {
	var y []int

	for _, val := range v {
		if val%2 == 0 {
			y = append(y, val)
		}
	}

	return f(y...) // Calling the function
}

func odd(f func(x ...int) int, v ...int) int {
	var y []int

	for _, val := range v {
		if val%2 != 0 {
			y = append(y, val)
		}
	}

	return f(y...)
}
