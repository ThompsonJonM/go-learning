package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7}
	fmt.Println(x) // [4 5 6 7]
	x = append(x, 8, 9, 10)
	fmt.Println(x) // [4 5 6 7 8 9 10]

	y := []int{1, 2, 3}
	fmt.Println(y) // [1 2 3]

	x = append(x, y...) // ... AFTER the slice tells Go to take all of the values
	fmt.Println(x)      // [4 5 6 7 8 9 10 1 2 3]
}
