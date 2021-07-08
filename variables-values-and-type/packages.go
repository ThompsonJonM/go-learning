package main

import "fmt"

func main() {
	// Println takes a variadic number of arguments
	// Returns the number in bytes and any errors with printing
	// To catch return values, assign the values to a variable
	num1, err := fmt.Println("Hello World", 42, true)
	fmt.Println(num1)
	fmt.Println(err)

	// Use underscores to toss away the returned value
	// Returned values MUST be assigned to something, whether it is a
	// variable (for use) or an underscore (for tossing)
	num2, _ := fmt.Println("This is another print line")
	fmt.Println(num2)
}
