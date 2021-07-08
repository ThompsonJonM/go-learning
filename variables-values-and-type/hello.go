package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	foo()
	fmt.Println("Printing something else now")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("and then we exited.")
}

// control flows:
// (1) sequential
// (2) loop; iterative
// (3) conditional
