package main

import "fmt"

func main() {
	// Use pointers when you don't want to pass around a lot of data, or when you
	// need to change the data at a specific location
	x := 0

	fmt.Println("X before:", &x)
	fmt.Println("X before:", x) // 0

	foo(&x)

	fmt.Println("X after:", &x)
	fmt.Println("X after:", x) // 43
}

func foo(y *int) {
	fmt.Println("Y before:", y) // 0
	fmt.Println("Y before:", *y)
	*y = 43
	fmt.Println("Y after:", y) // 43
	fmt.Println("Y after:", *y)
}
