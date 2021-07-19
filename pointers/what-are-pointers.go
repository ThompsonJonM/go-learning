package main

import "fmt"

func main() {
	a := 42

	fmt.Println(a)
	fmt.Println(&a) // Print address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) // Pointer to an int

	b := &a
	fmt.Println(b)
	fmt.Println(*b)  // Dereferencing: Prints the VALUE of the pointer
	fmt.Println(*&b) // Prints the VALUE and the address of the pointer

	*b = 43
	fmt.Println(a) // 43
}
