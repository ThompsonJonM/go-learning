package main

import "fmt"

func main() {
	x := []int{2, 3, 4, 5}
	fmt.Println(x[1])   // 3
	fmt.Println(x[1:])  // [3 4 5]
	fmt.Println(x[1:3]) // [3 4]

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}
