package main

import "fmt"

func main() {
	x := []int{2, 3, 4, 5}
	fmt.Println(len(x))

	for i, v := range x {
		fmt.Println(i, v)
	}
}
