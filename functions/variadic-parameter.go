package main

import "fmt"

func main() {
	foo(2, 3, 4)
}

func foo(x ...int) { // zero or more args can be passed
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
		fmt.Printf("We are adding %v with a total of %v\n", v, sum)
	}
}
