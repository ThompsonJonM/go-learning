package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6}
	x := foo(xi...) // Allows slice of int to be passed into this function
	fmt.Println("The total is", x)
}

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
		fmt.Printf("We are adding %v with a total of %v\n", v, sum)
	}

	return sum
}
