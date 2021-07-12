package main

import "fmt"

const (
	_  = iota             // Throw away 0 value
	kb = 1 << (iota * 10) // 10000000000
	mb = 1 << (iota * 10) // 100000000000000000000
	gb = 1 << (iota * 10) // 1000000000000000000000000000000
)

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x) // 4 100

	y := x << 1                    // Bit shift once to the left
	fmt.Printf("%d\t\t%b\n", y, y) // 8 1000

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
