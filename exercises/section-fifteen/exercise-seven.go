package main

import "fmt"

// Declaring a TYPE function outside of main
var f func() = func() {
	fmt.Println("Hello world from outside main")
}

func main() {
	f()

	f := func(n int) {
		total := 1
		for ; n > 0; n-- {
			total *= n
		}

		fmt.Println(total)
	}

	f(4)
}
