package main

import "fmt"

func main() {
	// if init; condition {}
	if x := 42; x == 42 {
		// x is limited to if block scope
		fmt.Println(x)
	}

	if true {
		fmt.Println("This is going to print.")
	}

	if false {
		fmt.Println("This will not.")
	}
}
