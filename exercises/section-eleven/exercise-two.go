package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	fmt.Println(x)

	for i, v := range x {
		fmt.Printf("The index is %v and value is %v\n", i, v)
	}

	fmt.Printf("%T", x)
}
