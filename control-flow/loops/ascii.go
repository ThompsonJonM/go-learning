package main

import "fmt"

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%+q\n", i, i) // Print value and ASCII char
	}
}
