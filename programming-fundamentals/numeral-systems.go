package main

import "fmt"

func main() {
	s := "H"

	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	n := bs[0]
	fmt.Println(n)         // Base 10
	fmt.Printf("%b\n", n)  // Binary
	fmt.Printf("%#X\n", n) // Hex
}
