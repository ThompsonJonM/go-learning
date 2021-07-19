package main

import "fmt"

func main() {
	x := []int{2, 3, 4}

	func() {
		fmt.Println("Hello world")
	}()

	n := func(n ...int) int {
		total := 0

		for _, v := range n {
			total += v
		}

		return total
	}(x...)

	fmt.Println(n)
}
