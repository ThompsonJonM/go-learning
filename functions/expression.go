package main

import "fmt"

func main() {
	f := func() { // Assign a function to a variable
		fmt.Println("Anonymous function ran")
	}

	f() // Call the variable to call the function
	fmt.Printf("%T\n", f)
}
