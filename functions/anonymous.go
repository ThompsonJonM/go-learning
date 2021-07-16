package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous function ran")
	}()

	func(x int) {
		fmt.Println("The meaning of life is:", x)
	}(42)
}
