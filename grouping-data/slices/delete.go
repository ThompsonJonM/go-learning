package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7}
	fmt.Println(x) // [4 5 6 7]

	x = append(x, 8, 9, 10)
	fmt.Println(x) // [4 5 6 7 8 9 10]

	x = append(x[:2], x[4:]...) // Delete 6 and 7 from slice
	fmt.Println(x)              // [4 5 8 9 10]
}
